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
	if err := m.Conn.Table("participation").Select("user_id", "room_id", "is_host").Create(&participation).Error; err != nil {
		return err
	}

	return nil
}
