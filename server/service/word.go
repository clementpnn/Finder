package service

import (
	"Finder/database/sqlc"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (s *WordService) InsertWord(domainID uuid.UUID) {
	counter := 1

	for {
		pages, err := s.pageRepo.GetPageByDomainId(domainID, counter)
		if err != nil {
			return
		}

		if len(*pages) == 0 {
			break
		}
		counter++

		for _, url := range *pages {
			s.processURL(url)
		}
	}
}

func (s *WordService) processURL(url sqlc.GetPageByDomainIdRow) {
	doc, err := getPage(url.Url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	text := doc.Text()
	wordCount := make(map[string]int32)
	words := strings.Fields(text)
	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	for word, count := range wordCount {
		wordId, err := s.wordRepo.InsertWord(word)
		if err == sql.ErrNoRows {
			wordId, err = s.wordRepo.GetWordId(word)
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}

		err = s.wordRepo.InsertWordPage(*wordId, url.ID, count)
		if err != nil {
			return
		}
	}
	err = s.pageRepo.UpdatePageWordCount(url.Url, int32(len(words)))
	if err != nil {
		return
	}
}
