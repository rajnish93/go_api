package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rajnish93/go_api/config"
	"github.com/rajnish93/go_api/db"
)

func init() {
  err := godotenv.Load(".env")
  if err != nil {
      log.Fatal("Error loading .env file")
  }
}

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World 👋!")
})
  db.Connection()
	port := config.ListenPort()
  app.Listen(port)
}