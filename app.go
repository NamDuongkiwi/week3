package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/NamDuongkiwi/Lab2/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	bookRouter := app.Group("/api/book")
	routes.ConfigBookRouter(&bookRouter) //http://localhost:3000/api/book
	reviewRouter := app.Group("/api/review")
	routes.ConfigReviewRouter(&reviewRouter) //http://localhost:3000/api/review
	app.Listen(":3000")
}