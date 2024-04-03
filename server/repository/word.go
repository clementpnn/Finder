package repository

import (
	"Finder/database/sqlc"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func (r *PostgresRepository) InsertWord(word string) (*uuid.UUID, error) {
	id, err := sqlc.New(r.db).InsertWord(context.Background(), word)
	if err == sql.ErrNoRows {
		return &id, err
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &id, nil
}

func (r *PostgresRepository) GetWordId(word string) (*uuid.UUID, error) {
	id, err := sqlc.New(r.db).GetWordId(context.Background(), word)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &id, nil
}

func (r *PostgresRepository) InsertWordPage(wordID uuid.UUID, pageID uuid.UUID, count int32) error {
	data := sqlc.InsertWordPageParams{
		WordID:    wordID,
		PageID:    pageID,
		WordCount: count,
	}
	err := sqlc.New(r.db).InsertWordPage(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
