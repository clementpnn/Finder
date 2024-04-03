package repository

import (
	"Finder/database/sqlc"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *PostgresRepository) IsExistDomain(domain string) (*bool, error) {
	exist, err := sqlc.New(r.db).IsExistDomain(context.Background(), domain)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &exist, nil
}

func (r *PostgresRepository) InsertDomain(domain string) (*uuid.UUID, error) {
	id, err := sqlc.New(r.db).InsertDomain(context.Background(), domain)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &id, nil
}

func (r *PostgresRepository) DeleteDomain(domain string) error {
	err := sqlc.New(r.db).DeleteDomain(context.Background(), domain)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
