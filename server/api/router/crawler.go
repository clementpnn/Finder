package router

import (
	"Finder/api/handler"

	"github.com/gofiber/fiber/v2"
)

func CrawlerRoutes(app *fiber.App, crawlerHandler *handler.CrawlerHandler) {
	crawlerGroup := app.Group("/crawler")
	crawlerGroup.Post("/index", crawlerHandler.Index)
	crawlerGroup.Post("/reindex", crawlerHandler.Reindex)
	crawlerGroup.Delete("/deindex", crawlerHandler.Deindex)
}
