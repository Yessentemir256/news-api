package handlers

import (
	"github.com/Yessentemir256/news-api/models"
	"github.com/go-reform/reform"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *reform.DB) {
	app.Post("/edit/:id", func(c *fiber.Ctx) error {
		return EditNewsHandler(c, db)
	})
	app.Get("/list", func(c *fiber.Ctx) error {
		return ListNewsHandler(c, db)
	})
}

func EditNewsHandler(c *fiber.Ctx, db *reform.DB) error {
	id := c.Params("id")
	news := new(models.News)
	if err := c.BodyParser(news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if _, err := db.FindByPrimaryKeyFrom("news", id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "News not found",
		})
	}
	if err := db.Update(news); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update news",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func ListNewsHandler(c *fiber.Ctx, db *reform.DB) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	var newsList []models.News
	if err := db.SelectAllFrom("news", "ORDER BY id ASC LIMIT ? OFFSET ?", limit, offset).Into(&newsList); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch news",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"news":    newsList,
	})
}
