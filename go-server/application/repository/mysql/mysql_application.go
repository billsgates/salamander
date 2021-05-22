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
