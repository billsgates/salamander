package mysql

import (
	"context"
	"go-server/domain"

	"gorm.io/gorm"
)

type mysqlRoundRepository struct {
	Conn *gorm.DB
}

func NewmysqlRoundRepository(Conn *gorm.DB) domain.RoundRepository {
	return &mysqlRoundRepository{Conn}
}

func (m *mysqlRoundRepository) AddRound(ctx context.Context, roomId int32, round *domain.Round) (err error) {
	if err := m.Conn.Table("rounds").Create(&round).Error; err != nil {
		return err
	}

	return nil
}
