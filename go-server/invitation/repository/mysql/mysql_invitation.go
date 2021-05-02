package mysql

import (
	"context"
	"go-server/domain"

	"gorm.io/gorm"
)

type mysqlInvitationRepository struct {
	Conn *gorm.DB
}

func NewmysqlInvitationRepository(Conn *gorm.DB) domain.InvitationRepository {
	return &mysqlInvitationRepository{Conn}
}

func (m *mysqlInvitationRepository) GenerateInvitationCode(ctx context.Context, roomId int32, code string) (err error) {
	invitation := domain.Invitation{RoomId: roomId, InvitationCode: code, IsValid: true}

	if err := m.Conn.Table("invitation_codes").Select("room_id", "invitation_code", "is_valid").Create(&invitation).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlInvitationRepository) ConsumeInvitationCode(ctx context.Context, code string) (roomId int32, err error) {
	var invitation *domain.Invitation

	if err := m.Conn.Table("invitation_codes").Where("invitation_code = ?", code).First(&invitation).Error; err != nil {
		return -1, err
	}

	if err := m.Conn.Table("invitation_codes").Where("invitation_code = ?", code).Update("is_valid", false).Error; err != nil {
		return -1, err
	}

	return invitation.RoomId, nil
}
