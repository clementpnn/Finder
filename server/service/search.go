package service

import (
	"Finder/database/sqlc"
	"Finder/domain/model"
	"encoding/json"
	"strings"
)

func (s *SearchService) Search(input string, page int32) (*[]model.SearchOutput, *float64, error) {
	words := strings.Fields(strings.ToLower(input))
	databasePage := page - 1
	data, err := s.searchRepo.Search(words, databasePage)
	if err != nil {
		return nil, nil, err
	}

	newData, totalPages, err := processData(*data)
	if err != nil {
		return nil, nil, err
	}

	return newData, totalPages, nil
}

func processData(data []sqlc.SearchRow) (*[]model.SearchOutput, *float64, error) {
	newDataArray := make([]model.SearchOutput, 0)
	for _, d := range data {
		var metadata []model.MetaData
		if err := json.Unmarshal([]byte(d.MetaData), &metadata); err != nil {
			return nil, nil, err
		}

		findMetaDataContent := func(metadata []model.MetaData, names []string) string {
			for _, m := range metadata {
				for _, name := range names {
					if m.Name == name || m.Property == name {
						return m.Content
					}
				}
			}
			return ""
		}

		descriptionKeys := []string{"description", "og:description"}
		imageKeys := []string{"image", "og:image"}

		newDataArray = append(newDataArray, model.SearchOutput{
			Url:         d.Url,
			Title:       d.Title,
			Description: findMetaDataContent(metadata, descriptionKeys),
			Image:       findMetaDataContent(metadata, imageKeys),
		})
	}

	totalPages := data[0].TotalPages
	if totalPages > 10 {
		totalPages = 10
	}

	return &newDataArray, &totalPages, nil
}
