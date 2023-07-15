package main

import (
	// "fmt"

	"fmt"
	useful "github.com/verb5/blog/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		body := fmt.Sprintf(useful.)
		return c.SendString()

	})
	app.Listen(":3000")
}
