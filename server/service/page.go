package service

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
)

func (s *PageService) FecthRobotsTxt(domain string) *[]string {
	robotsURL := fmt.Sprintf("%s/robots.txt", domain)

	resp, err := http.Get(robotsURL)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var disallowedPaths []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Disallow:") {
			path := strings.TrimSpace(strings.TrimPrefix(line, "Disallow:"))
			if path != "" {
				disallowedPaths = append(disallowedPaths, path)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &disallowedPaths
}

func (s *PageService) InsertPage(domain string, disallowedPaths []string, domainID uuid.UUID) {
	var wg sync.WaitGroup
	urls := make(chan string, 100)
	s.pageRepo.InsertPage(domain, domainID)
	urls <- domain

	go func() {
		wg.Wait()
		close(urls)
	}()

	for url := range urls {
		wg.Add(1)
		go s.processURL(url, domain, domainID, disallowedPaths, urls, &wg)
	}
}

func isDisallowedPaths(urlToCheck string, disallowedPaths []string) bool {
	for _, path := range disallowedPaths {
		if strings.Contains(urlToCheck, path) {
			return true
		}
	}
	return false
}

func getPage(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("status code is not 200 OK")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func getPageData(doc *goquery.Document) (string, json.RawMessage, error) {
	var metaAttributesSlice []map[string]string

	title := doc.Find("title").First().Text()
	if title == "" {
		title = "No title"
	}
	doc.Find("meta").Each(func(i int, g *goquery.Selection) {
		metaAttributes := make(map[string]string)
		for _, attr := range g.Nodes[0].Attr {
			metaAttributes[attr.Key] = attr.Val
		}
		metaAttributesSlice = append(metaAttributesSlice, metaAttributes)
	})

	metaDataJSON, err := json.Marshal(metaAttributesSlice)
	if err != nil {
		return "", nil, err
	}
	return title, metaDataJSON, nil
}

func resolveUrl(baseUrl, foundUrl string) (string, error) {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	found, err := url.Parse(foundUrl)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(found).String(), nil
}

func checkHost(domain string, absoluteUrl string) error {
	parsedURL, err := url.Parse(absoluteUrl)
	if err != nil {
		return err
	}
	parsedDomain, err := url.Parse(domain)
	if err != nil {
		return err
	}
	if parsedURL.Host != parsedDomain.Host {
		return errors.New("URL is not from the same domain")
	}
	return nil
}

func (s *PageService) processURL(urlToCheck string, domain string, domainID uuid.UUID, disallowedPaths []string, urls chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	doc, err := getPage(urlToCheck)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	title, metaDataJSON, err := getPageData(doc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	urlToCheckId, err := s.pageRepo.UpdatePageData(urlToCheck, title, metaDataJSON)
	if err != nil {
		return
	}

	doc.Find("a").Each(func(i int, g *goquery.Selection) {
		foundUrl, exists := g.Attr("href")
		if !exists {
			return
		}
		if strings.Contains(foundUrl, "#") {
			return
		}

		foundUrl = strings.TrimSuffix(foundUrl, "/")
		absoluteUrl, err := resolveUrl(domain, foundUrl)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = checkHost(domain, absoluteUrl)
		if err != nil {
			return
		}

		_, err = getPage(absoluteUrl)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if isDisallowedPaths(urlToCheck, disallowedPaths) {
			return
		}

		exist, err := s.pageRepo.IsExistPage(absoluteUrl)
		if err != nil {
			return
		}

		if *exist {
			absoluteUrlId, err := s.pageRepo.GetPageIdByUrl(absoluteUrl)
			if err != nil {
				return
			}
			s.pageRepo.InsertPageReferral(*urlToCheckId, *absoluteUrlId)
			return
		}

		s.pageRepo.InsertPage(absoluteUrl, domainID)
		urls <- absoluteUrl
	})
}
