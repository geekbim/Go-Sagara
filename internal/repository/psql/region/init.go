package postgres_repository

import (
	"database/sql"
	"sagara/domain/repository"
)

type regionRepository struct {
	db *sql.DB
}

func NewRegionRepository(db *sql.DB) repository.RegionRepository {
	return &regionRepository{
		db: db,
	}
}
