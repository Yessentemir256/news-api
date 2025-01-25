// main.go
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
	db := database.InitDatabase()
	defer db.Close()

	// Create Fiber app
	app := fiber.New()
	app.Use(logger.New())

	// Define routes
	handlers.SetupRoutes(app, db)

	// Start server
	log.Fatal(app.Listen(config.GetServerAddress()))
}
