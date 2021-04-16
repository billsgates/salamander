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

func (m *mysqlUserRepository) Create(ctx context.Context, user *domain.User) (res *domain.User, err error) {
	m.Conn.Create(&user)

	return user, nil
}

func (m *mysqlUserRepository) FetchAll(ctx context.Context) (res []domain.User, err error) {
	var users []domain.User
	m.Conn.Find(&users)

	return users, nil
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id string) (res *domain.User, err error) {
	var user *domain.User
	m.Conn.First(&user, id)

	return user, nil
}

func (m *mysqlUserRepository) GetByEmailPassword(ctx context.Context, email string, password string) (res *domain.User, err error) {
	var user *domain.User

	if err := m.Conn.Where("email = ? AND password_digest = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
