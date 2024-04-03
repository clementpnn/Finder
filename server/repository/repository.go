package repository

import (
	"Finder/domain/facade"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewDomainRepository(db *sql.DB) facade.DomainRepository {
	return &PostgresRepository{db}
}

func NewPageRepository(db *sql.DB) facade.PageRepository {
	return &PostgresRepository{db}
}

func NewWordRepository(db *sql.DB) facade.WordRepository {
	return &PostgresRepository{db}
}

func NewSearchRepository(db *sql.DB) facade.SearchRepository {
	return &PostgresRepository{db}
}
