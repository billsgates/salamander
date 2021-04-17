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
	m.Conn.Create(&room)

	return room, nil
}
