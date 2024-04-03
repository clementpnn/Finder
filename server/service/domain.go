package service

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

func (s *DomainService) IsValidDomain(inputUrl string) (*string, error) {
	if !strings.HasPrefix(inputUrl, "http://") && !strings.HasPrefix(inputUrl, "https://") {
		inputUrl = "https://" + inputUrl
	}

	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	domainURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	_, err = getPage(domainURL)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &domainURL, nil
}

func (s *DomainService) IsExistDomain(domain string) (*bool, error) {
	return s.domainRepo.IsExistDomain(domain)
}

func (s *DomainService) InsertDomain(domain string) (*uuid.UUID, error) {
	return s.domainRepo.InsertDomain(domain)
}

func (s *DomainService) DeleteDomain(domain string) error {
	return s.domainRepo.DeleteDomain(domain)
}
