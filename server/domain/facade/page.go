package facade

import (
	"Finder/database/sqlc"
	"encoding/json"

	"github.com/google/uuid"
)

type PageService interface {
	FecthRobotsTxt(domain string) *[]string
	InsertPage(domain string, disallowedPaths []string, domainID uuid.UUID)
}

type PageRepository interface {
	IsExistPage(url string) (*bool, error)
	GetPageIdByUrl(url string) (*uuid.UUID, error)
	InsertPage(url string, domainID uuid.UUID) error
	UpdatePageData(url string, title string, metaData json.RawMessage) (*uuid.UUID, error)
	InsertPageReferral(urlId uuid.UUID, referralID uuid.UUID) error
	UpdatePageWordCount(url string, wordCount int32) error
	GetPageByDomainId(domainID uuid.UUID, index int) (*[]sqlc.GetPageByDomainIdRow, error)
}
