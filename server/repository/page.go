package repository

import (
	"Finder/database/sqlc"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func (r *PostgresRepository) IsExistPage(url string) (*bool, error) {
	exist, err := sqlc.New(r.db).IsExistPage(context.Background(), url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &exist, nil
}

func (r *PostgresRepository) GetPageIdByUrl(url string) (*uuid.UUID, error) {
	id, err := sqlc.New(r.db).GetPageIdByUrl(context.Background(), url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &id, nil
}

func (r *PostgresRepository) InsertPage(url string, domainID uuid.UUID) error {
	data := sqlc.InsertPageParams{
		Url:      url,
		DomainID: domainID,
	}

	err := sqlc.New(r.db).InsertPage(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *PostgresRepository) UpdatePageData(url string, title string, metaData json.RawMessage) (*uuid.UUID, error) {
	data := sqlc.UpdatePageDataParams{
		Url:      url,
		Title:    title,
		MetaData: metaData,
	}

	id, err := sqlc.New(r.db).UpdatePageData(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &id, err
}

func (r *PostgresRepository) InsertPageReferral(urlId uuid.UUID, referralID uuid.UUID) error {
	data := sqlc.InsertPageReferralParams{
		PageID:     urlId,
		ReferralID: referralID,
	}
	err := sqlc.New(r.db).InsertPageReferral(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *PostgresRepository) UpdatePageWordCount(url string, wordCount int32) error {
	data := sqlc.UpdatePageWordCountParams{
		Url:       url,
		WordCount: wordCount,
	}
	err := sqlc.New(r.db).UpdatePageWordCount(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *PostgresRepository) GetPageByDomainId(domainID uuid.UUID, index int) (*[]sqlc.GetPageByDomainIdRow, error) {
	data := sqlc.GetPageByDomainIdParams{
		DomainID: domainID,
		Column2:  index,
	}
	pages, err := sqlc.New(r.db).GetPageByDomainId(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &pages, nil
}
