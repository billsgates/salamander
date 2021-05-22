package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"go-server/auth"
	"go-server/domain"
	adapterqueue "go-server/internal/adapter/queue"
	helper "go-server/internal/common"
	"go-server/internal/infrastructure/queue"

	"github.com/dgrijalva/jwt-go/v4"
)

type authUsecase struct {
	userRepo       domain.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
	producer       adapterqueue.Producer
}

func NewAuthUseCase(
	userRepo domain.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration,
	queue *queue.RabbitMQHandler) domain.AuthUsecase {
	return &authUsecase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
		producer:       adapterqueue.NewProducer(queue.Channel(), "verification"),
	}
}

func (a *authUsecase) SignUp(ctx context.Context, name string, email string, password string) (res string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &domain.User{
		Name:           name,
		Email:          email,
		PasswordDigest: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	err = a.userRepo.Create(ctx, user)
	if err != nil {
		return "", err
	}

	message := helper.EncodeToBytes(user)
	message = helper.Compress(message)
	a.producer.Publish(message)
	return a.SignIn(ctx, email, password)
}

func (a *authUsecase) SignIn(ctx context.Context, email string, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetByEmailPassword(ctx, email, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := domain.AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *authUsecase) ParseToken(ctx context.Context, accessToken string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &domain.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
