package repository

import (
	"database/sql"

	"github.com/febryanmz/Data-Diri/features/profile"
)

type profileRepository struct {
	db *sql.DB
}

func NewRaw(db *sql.DB) profile.RepositoryInterface {
	return &profileRepository{
		db: db,
	}
}
