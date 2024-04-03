package facade

import (
	"Finder/database/sqlc"
	"Finder/domain/model"
)

type SearchService interface {
	Search(input string, page int32) (*[]model.SearchOutput, *float64, error)
}

type SearchRepository interface {
	Search(input []string, page int32) (*[]sqlc.SearchRow, error)
}
