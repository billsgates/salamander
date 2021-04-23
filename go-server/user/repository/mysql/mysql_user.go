package mysql

import (
	"context"
	"go-server/domain"

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
