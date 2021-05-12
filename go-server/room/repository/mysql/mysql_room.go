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

func (m *mysqlRoomRepository) Create(ctx context.Context, room *domain.Room) (roomId int32, err error) {
	if err := m.Conn.Select("max_count", "admin_id", "service_id", "plan_name", "payment_period", "is_public").Create(&room).Error; err != nil {
		return -1, err
	}

	return room.Id, nil
}

func (m *mysqlRoomRepository) Update(ctx context.Context, roomId int32, room *domain.Room) (err error) {
	if err := m.Conn.Table("rooms").Select("max_count", "service_id", "plan_name", "payment_period", "is_public").Where("room_id = ?", roomId).Updates(&room).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlRoomRepository) Delete(ctx context.Context, roomId int32) (err error) {
	var room *domain.Room
	if err := m.Conn.Table("rooms").Where("room_id = ?", roomId).Delete(&room).Error; err != nil {
		return err
	}

	return nil
}
