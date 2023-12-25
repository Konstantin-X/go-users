package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

const defaultPort string = "3000"

func Setup() {
	app := fiber.New(fiber.Config{})
	api := fiber.New(fiber.Config{})

	app.Mount("/api", api)
	app.Use(logger.New())

	api.Route("/users", func(router fiber.Router) {
		router.Delete("", DeleteUser)
		router.Get("", ListUsers)
		router.Get("/{id}", ReadUser)
		router.Post("", CreateUser)
		router.Put("/{id}", UpdateUser)
	})

	api.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "All systems go!",
		})
	})

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultPort
	}

	log.Fatal(app.Listen("localhost:" + port))
}
