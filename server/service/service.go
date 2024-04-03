package service

import (
	"Finder/domain/facade"
)

type DomainService struct {
	domainRepo facade.DomainRepository
}

type PageService struct {
	pageRepo facade.PageRepository
}

type WordService struct {
	wordRepo facade.WordRepository
	pageRepo facade.PageRepository
}

type SearchService struct {
	searchRepo facade.SearchRepository
}

func NewDomainService(domainRepo facade.DomainRepository) *DomainService {
	return &DomainService{domainRepo}
}

func NewPageService(pageRepo facade.PageRepository) *PageService {
	return &PageService{pageRepo}
}

func NewWordService(wordRepo facade.WordRepository, pageRepo facade.PageRepository) *WordService {
	return &WordService{wordRepo, pageRepo}
}

func NewSearchService(searchRepo facade.SearchRepository) *SearchService {
	return &SearchService{searchRepo}
}
