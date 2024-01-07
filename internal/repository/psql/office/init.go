package postgres_repository

import (
	"database/sql"
	"sagara/domain/repository"
)

type officeRepository struct {
	db *sql.DB
}

func NewOfficeRepository(db *sql.DB) repository.OfficeRepository {
	return &officeRepository{
		db: db,
	}
}
