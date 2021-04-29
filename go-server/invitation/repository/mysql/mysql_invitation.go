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
