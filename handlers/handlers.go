package handlers

import (
	"github.com/Yessentemir256/news-api/models"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/reform.v1"
	"strconv"
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
	newsID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Поиск новости
	news := &models.News{}
	if err := db.FindByPrimaryKeyTo(news, newsID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "News not found",
		})
	}

	// Обновление новости
	input := new(models.News)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if input.Title != "" {
		news.Title = input.Title
	}
	if input.Content != "" {
		news.Content = input.Content
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

	// Запрос новостей
	records, err := db.SelectAllFrom(models.NewsTable, "ORDER BY id ASC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch news",
		})
	}

	// Преобразование записей в срез News
	newsList := make([]models.News, len(records))
	for i, record := range records {
		newsList[i] = *record.(*models.News)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"news":    newsList,
	})
}
