package mysql

import (
	"context"
	"go-server/domain"

	"gorm.io/gorm"
)

type mysqlParticipationRepository struct {
	Conn *gorm.DB
}

func NewmysqlParticipationRepository(Conn *gorm.DB) domain.ParticipationRepository {
	return &mysqlParticipationRepository{Conn}
}

func (m *mysqlParticipationRepository) Create(ctx context.Context, participation *domain.Participation) (err error) {
	if err := m.Conn.Table("participation").Select("user_id", "room_id", "payment_status", "is_host").Create(&participation).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlParticipationRepository) GetJoinedRooms(ctx context.Context, id int32) (res []domain.RoomInfo, err error) {
	var rooms []domain.RoomInfo
	if err := m.Conn.Table("participation").Select("service_providers.name, rooms.plan_name, participation.is_host").Joins("JOIN rooms ON rooms.room_id = participation.room_id").Joins("JOIN service_providers ON service_providers.id = rooms.service_id").Where("participation.user_id = ?", id).Scan(&rooms).Error; err != nil {
		return nil, err
	}

	return rooms, nil
}

func (m *mysqlParticipationRepository) IsAdmin(ctx context.Context, roomId int32, userId int32) (res bool, err error) {
	var participation domain.Participation
	if err := m.Conn.Table("participation").Where("room_id = ? AND user_id = ?", roomId, userId).First(&participation).Error; err != nil {
		return false, err
	}

	return participation.IsHost, nil
}
