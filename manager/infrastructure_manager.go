package manager

import "github.com/jmoiron/sqlx"

type infrastructureManager struct {
	sqlDB *sqlx.DB
}

func (m infrastructureManager) GetDB() *sqlx.DB {
	return m.sqlDB
}
