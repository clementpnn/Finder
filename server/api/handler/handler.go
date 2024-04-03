package handler

import "Finder/domain/facade"

type CrawlerHandler struct {
	domainService facade.DomainService
	pageService   facade.PageService
	wordService   facade.WordService
}

type SearchHandler struct {
	searchService facade.SearchService
}

func NewCrawlerHandler(domainService facade.DomainService, pageService facade.PageService, wordService facade.WordService) *CrawlerHandler {
	return &CrawlerHandler{domainService, pageService, wordService}
}

func NewSearchHandler(searchService facade.SearchService) *SearchHandler {
	return &SearchHandler{searchService}
}
