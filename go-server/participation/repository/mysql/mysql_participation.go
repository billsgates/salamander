package mysql

import (
	"context"
	"go-server/domain"
	"time"

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

func (m *mysqlParticipationRepository) GetJoinedRooms(ctx context.Context, id int32) (res []domain.RoomItem, err error) {
	var rooms []domain.RoomItem
	if err := m.Conn.Table("participation").Select("service_providers.name as service_name, rooms.room_id, rooms.plan_name, rooms.room_status, participation.is_host, participation.payment_status, plans.cost").
		Joins("JOIN rooms ON rooms.room_id = participation.room_id").
		Joins("JOIN service_providers ON service_providers.id = rooms.service_id").
		Joins("JOIN plans ON plans.plan_name = rooms.plan_name AND plans.service_id = rooms.service_id").
		Where("participation.user_id = ?", id).Scan(&rooms).Error; err != nil {
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

func (m *mysqlParticipationRepository) LeaveRoom(ctx context.Context, roomId int32, userId int32) (err error) {
	var participation domain.Participation
	if err := m.Conn.Table("participation").Where("room_id = ? AND user_id = ?", roomId, userId).Delete(&participation).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlParticipationRepository) GetRoomInfo(c context.Context, roomId int32) (res *domain.RoomInfoResponse, err error) {
	var roomInfo *domain.RoomInfoResponse
	if err := m.Conn.Table("rooms").Select("service_providers.name as service_name, rooms.room_id, rooms.is_public, rooms.announcement, rooms.max_count, rooms.plan_name, rooms.room_status, users.name as admin_name, users.email as admin_email, users.rating as admin_rating, users.phone as admin_phone, plans.cost as payment_fee").
		Joins("JOIN users ON users.id = rooms.admin_id").
		Joins("JOIN plans ON plans.plan_name = rooms.plan_name AND plans.service_id = rooms.service_id").
		Joins("JOIN service_providers ON service_providers.id = plans.service_id").
		Where("rooms.room_id = ?", roomId).First(&roomInfo).Error; err != nil {
		return nil, err
	}

	return roomInfo, nil
}

func (m *mysqlParticipationRepository) GetRoomAdmin(c context.Context, roomId int32) (res *domain.User, err error) {
	var admin *domain.User
	if err := m.Conn.Table("participation").Select("users.name AS name, users.email AS email, users.rating AS rating, users.phone AS phone").
		Joins("JOIN rooms ON rooms.room_id = participation.room_id").
		Joins("JOIN users ON users.id = rooms.admin_id").
		Where("participation.room_id = ?", roomId).Find(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func (m *mysqlParticipationRepository) GetRoomMembers(c context.Context, roomId int32) (res []domain.Participation, err error) {
	var members []domain.Participation
	if err := m.Conn.Table("participation").Select("users.id AS user_id, users.name AS user_name, participation.payment_status").
		Joins("JOIN users ON users.id = participation.user_id").
		Where("participation.room_id = ?", roomId).Scan(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}

func (m *mysqlParticipationRepository) GetRoomMemberByStartingTime(c context.Context, starting_time time.Time) (res []domain.Participation, err error) {
	var members []domain.Participation
	if err := m.Conn.Table("participation").Select("users.id AS user_id, users.name AS user_name, participation.payment_status").
		Joins("JOIN users ON users.id = participation.user_id").
		Joins("JOIN rooms ON rooms.room_id = participation.room_id").
		Where("rooms.starting_time = ?", starting_time).Scan(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}
