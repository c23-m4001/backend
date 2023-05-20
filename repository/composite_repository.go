package repository

import (
	"github.com/jmoiron/sqlx"
)

type CompositeRepository interface {
}

type compositeRepository struct {
	db *sqlx.DB
}

func NewCompositeRepository(
	db *sqlx.DB,
) CompositeRepository {
	return &compositeRepository{
		db: db,
	}
}
