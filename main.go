package main

import (
	// "fmt"

	"fmt"
	useful "github.com/verb5/blog/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		body := fmt.Sprintf("%v\n%v\n",useful.GetHostname(),useful.GetIPAddress())
		return c.SendString(body)

	})
	app.Listen(":3000")
}
