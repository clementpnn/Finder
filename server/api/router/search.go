package router

import (
	"Finder/api/handler"

	"github.com/gofiber/fiber/v2"
)

func SearchRoutes(app *fiber.App, searchHandler *handler.SearchHandler) {
	searchGroup := app.Group("/search")
	searchGroup.Post("/", searchHandler.Search)
}
