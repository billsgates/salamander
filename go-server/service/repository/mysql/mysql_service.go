package mysql

import (
	"context"
	"go-server/domain"

	"gorm.io/gorm"
)

type mysqlServiceRepository struct {
	Conn *gorm.DB
}

func NewmysqlServiceRepository(Conn *gorm.DB) domain.ServiceRepository {
	return &mysqlServiceRepository{Conn}
}

func (m *mysqlServiceRepository) FetchAll(ctx context.Context) (res []domain.Service, err error) {
	var services []domain.Service
	m.Conn.Table("service_providers").Find(&services)
	// m.Conn.Table("service_providers").Select("service_providers.name, plans.plan_name").Joins("left join plans on plans.service_id = service_providers.id").Scan(&services)
	return services, nil
}
