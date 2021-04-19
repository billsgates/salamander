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

func (m *mysqlParticipationRepository) Create(ctx context.Context, participation *domain.Participation) (res *domain.Participation, err error) {
	m.Conn.Table("participation").Select("user_id", "room_id", "is_host").Create(&participation)

	return participation, nil
}
