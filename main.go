package main

import (
	"log"

	"github.com/Yessentemir256/news-api/config"
	"github.com/Yessentemir256/news-api/database"
	"github.com/Yessentemir256/news-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize configuration
	config.InitConfig()

	// Initialize database
	db, sqlDB := database.InitDatabase()
	defer sqlDB.Close() // Close the underlying *sql.DB

	// Create Fiber app
	app := fiber.New()
	app.Use(logger.New())

	// Define routes
	handlers.SetupRoutes(app, db)

	// Start server
	log.Fatal(app.Listen(config.GetServerAddress()))
}
