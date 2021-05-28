package mysql

import (
	"context"
	"go-server/domain"

	"gorm.io/gorm"
)

type mysqlApplicationRepository struct {
	Conn *gorm.DB
}

func NewmysqlApplicationRepository(Conn *gorm.DB) domain.ApplicationRepository {
	return &mysqlApplicationRepository{Conn}
}

func (m *mysqlApplicationRepository) Create(ctx context.Context, roomId int32, userId int32) (err error) {
	application := &domain.Application{
		RoomId: roomId,
		UserId: userId,
	}
	if err := m.Conn.Table("applications").Select("room_id", "user_id").Create(&application).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlApplicationRepository) FetchAll(ctx context.Context, roomId int32) (res []domain.Application, err error) {
	var applications []domain.Application

	if err := m.Conn.Table("applications").Select("applications.created_at as application_date, users.id as user_id, users.name as user_name, users.rating as user_rating").
		Joins("JOIN users ON users.id = applications.user_id").
		Where("room_id = ?", roomId).Scan(&applications).Error; err != nil {
		return nil, err
	}

	return applications, nil
}
