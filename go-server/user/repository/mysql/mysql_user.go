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

func (m *mysqlUserRepository) Fetch(ctx context.Context) (res []domain.User, err error) {
	var users []domain.User
	m.Conn.Find(&users)

	return users, nil
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id string) (res domain.User, err error) {
	var user domain.User
	m.Conn.First(&user, id)

	return user, nil
}
