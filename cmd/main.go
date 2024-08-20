package main

import (
	"os"

	"RankEdge/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	utils.ConnectToRedis()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(os.Getenv("HTTP_PORT"))
}
