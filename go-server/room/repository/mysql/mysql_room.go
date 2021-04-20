package mysql

import (
	"context"
	"go-server/domain"

	"gorm.io/gorm"
)

type mysqlRoomRepository struct {
	Conn *gorm.DB
}

func NewmysqlRoomRepository(Conn *gorm.DB) domain.RoomRepository {
	return &mysqlRoomRepository{Conn}
}

func (m *mysqlRoomRepository) Create(ctx context.Context, room *domain.Room) (res *domain.Room, err error) {
	if err := m.Conn.Select("max_count", "admin_id", "service_id", "plan_name").Create(&room).Error; err != nil {
		return nil, err
	}

	return room, nil
}
