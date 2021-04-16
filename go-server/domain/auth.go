package domain

import (
	"context"

	"github.com/dgrijalva/jwt-go/v4"
)

const CtxUserKey = "user"

type AuthClaims struct {
	jwt.StandardClaims
	User *User `json:"user"`
}

type AuthUsecase interface {
	SignUp(ctx context.Context, name, email, password string) (*User, error)
	SignIn(ctx context.Context, email, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*User, error)
}
