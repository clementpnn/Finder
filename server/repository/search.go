package repository

import (
	"Finder/database/sqlc"
	"context"
	"fmt"
)

func (r *PostgresRepository) Search(input []string, page int32) (*[]sqlc.SearchRow, error) {
	data := sqlc.SearchParams{
		Column1: input,
		Offset:  page * 20,
	}
	urls, err := sqlc.New(r.db).Search(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &urls, nil
}
