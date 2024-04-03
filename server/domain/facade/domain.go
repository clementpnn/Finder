package facade

import (
	"github.com/google/uuid"
)

type DomainService interface {
	IsValidDomain(inputUrl string) (*string, error)
	IsExistDomain(domain string) (*bool, error)
	InsertDomain(domain string) (*uuid.UUID, error)
	DeleteDomain(domain string) error
}

type DomainRepository interface {
	IsExistDomain(domain string) (*bool, error)
	InsertDomain(domain string) (*uuid.UUID, error)
	DeleteDomain(domain string) error
}
