package handler

import (
	"Finder/domain/model"

	"github.com/gofiber/fiber/v2"
)

// Index handler is used to index a domain.
// @Summary Index a domain
// @Description Index a domain, insert its pages and words into the database
// @Tags Crawler
// @Accept json
// @Produce json
// @Param input body model.Url true "URL of the domain to index"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /crawler/index [post]
func (h *CrawlerHandler) Index(c *fiber.Ctx) error {
	input := &model.Url{}
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	domain, err := h.domainService.IsValidDomain(input.URL)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	exist, err := h.domainService.IsExistDomain(*domain)
	if *exist || err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	domainId, error := h.domainService.InsertDomain(*domain)
	if error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	disallowedPaths := h.pageService.FecthRobotsTxt(*domain)

	h.pageService.InsertPage(*domain, *disallowedPaths, *domainId)

	h.wordService.InsertWord(*domainId)

	return c.SendStatus(fiber.StatusOK)
}

// Reindex handler is used to reindex a domain.
// @Summary Reindex a domain
// @Description Reindex a domain, delete its existing data and reinsert pages and words into the database
// @Tags Crawler
// @Accept json
// @Produce json
// @Param input body model.Url true "URL of the domain to reindex"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /crawler/reindex [post]
func (h *CrawlerHandler) Reindex(c *fiber.Ctx) error {
	input := &model.Url{}
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	domain, err := h.domainService.IsValidDomain(input.URL)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	exist, err := h.domainService.IsExistDomain(*domain)
	if !*exist || err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if error := h.domainService.DeleteDomain(*domain); error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	exist, err = h.domainService.IsExistDomain(*domain)
	if *exist || err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	domainId, error := h.domainService.InsertDomain(*domain)
	if error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	disallowedPaths := h.pageService.FecthRobotsTxt(*domain)

	h.pageService.InsertPage(*domain, *disallowedPaths, *domainId)

	h.wordService.InsertWord(*domainId)

	return c.SendStatus(fiber.StatusOK)
}

// Deindex handler is used to deindex a domain.
// @Summary Deindex a domain
// @Description Deindex a domain, delete its existing data from the database
// @Tags Crawler
// @Accept json
// @Produce json
// @Param input body model.Url true "URL of the domain to deindex"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /crawler/deindex [post]
func (h *CrawlerHandler) Deindex(c *fiber.Ctx) error {
	input := &model.Url{}
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	domain, err := h.domainService.IsValidDomain(input.URL)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	exist, err := h.domainService.IsExistDomain(*domain)
	if !*exist || err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if error := h.domainService.DeleteDomain(*domain); error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
