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

func (m *mysqlUserRepository) GetByID(ctx context.Context, id string) (res *domain.UserInfo, err error) {
	var user *domain.UserInfo
	if err := m.Conn.Table("users").Where("users.id = ?", id).First(&user).Error; err != nil {
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
	if err := m.Conn.Table("users").
		Where("id = ?", user.Id).
		Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) UpdateRating(ctx context.Context, id string, rating float32) (err error) {
	if err := m.Conn.Table("users").
		Where("id = ?", id).
		Update("rating", rating).
		Update("rating_count", gorm.Expr("rating_count + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}
