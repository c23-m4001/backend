package manager

import (
	"capstone/config"
	"capstone/infrastructure"

	"github.com/jmoiron/sqlx"
)

type InfrastructureManager interface {
	GetDB() *sqlx.DB
}

type infrastructureManager struct {
	sqlDB *sqlx.DB
}

func (m infrastructureManager) GetDB() *sqlx.DB {
	return m.sqlDB
}

func newInfrastructureManager() InfrastructureManager {
	return &infrastructureManager{
		sqlDB: infrastructure.NewPostgreSqlDB(config.GetPostgresConfig()),
	}
}
