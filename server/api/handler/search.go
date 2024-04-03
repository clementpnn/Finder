package handler

import (
	"Finder/domain/model"

	"github.com/gofiber/fiber/v2"
)

func (h *SearchHandler) Search(c *fiber.Ctx) error {
	input := &model.SearchInput{}
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if input.Page < 1 || input.Page > 10 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	urls, totalPages, err := h.searchService.Search(input.Input, input.Page)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if urls == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if *totalPages == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if int32(*totalPages) < input.Page {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{"pages": *urls, "totalPages": int(*totalPages)})
}
