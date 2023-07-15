package main

import (
	"github.com/gofiber/fiber/v2"
	u "github.com/verb5/blog/utils"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		body := u.SystemInfo{
			Hostname:  u.GetHostname(),
			IPAddress: u.GetIPAddress(),
		}

		return c.JSON(body)

	})
	app.Listen(":3000")
}
