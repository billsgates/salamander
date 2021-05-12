package mysql

import (
	"context"
	"go-server/domain"
	"time"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewmysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{Conn}
}

func (m *mysqlUserRepository) Create(ctx context.Context, user *domain.User) (err error) {
	if err := m.Conn.Select("name", "email", "password_digest").Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) FetchAll(ctx context.Context) (res []domain.User, err error) {
	var users []domain.User
	if err := m.Conn.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id string) (res *domain.User, err error) {
	var user *domain.User
	if err := m.Conn.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *mysqlUserRepository) GetByEmailPassword(ctx context.Context, email string, password string) (res *domain.User, err error) {
	var user *domain.User

	if err := m.Conn.Where("email = ? AND password_digest = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, user *domain.UserRequest) (err error) {
	if user.ImageUrl != "" {
		if err := m.Conn.Table("users").Where("id = ?", user.Id).
			Update("name", user.Name).Update("email", user.Email).
			Update("image_url", user.ImageUrl).
			Update("updated_at", time.Now()).
			Error; err != nil {
			return err
		}
		return nil
	}
	if err := m.Conn.Table("users").Where("id = ?", user.Id).
		Update("name", user.Name).
		Update("email", user.Email).
		Update("updated_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
