package facade

import "github.com/google/uuid"

type WordService interface {
	InsertWord(domainID uuid.UUID)
}

type WordRepository interface {
	InsertWord(word string) (*uuid.UUID, error)
	GetWordId(word string) (*uuid.UUID, error)
	InsertWordPage(wordID uuid.UUID, pageID uuid.UUID, count int32) error
}
