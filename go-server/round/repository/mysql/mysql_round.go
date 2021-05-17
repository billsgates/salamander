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

func (m *mysqlRoundRepository) GetRound(ctx context.Context, roomId int32) (res *domain.RoundInfo, err error) {
	var roundInfo *domain.RoundInfo

	if err := m.Conn.Table("rounds").Where("room_id = ?", roomId).First(&roundInfo).Error; err != nil {
		return roundInfo, err
	}

	return roundInfo, nil
}
