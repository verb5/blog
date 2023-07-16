package main

import (
	"github.com/gofiber/fiber/v2"
	u "github.com/verb5/blog/utils"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {

		conf := u.ParseYaml("config.yaml")
		body := u.SystemInfo{
			Hostname:  u.GetHostname(),
			IPAddress: u.GetIPAddress(),
			Location:  conf.ServerInfo.Location,
			Year:      conf.ServerInfo.Year,
			Email:     conf.ServerInfo.Email,
		}

		return c.JSON(body)

	})
	app.Listen(":3000")
}
