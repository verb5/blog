package main

import (
	// "fmt"

	"fmt"
	useful "github.com/verb5/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	app := fiber.New()
	app.Get("/:message", func(c *fiber.Ctx) error {
		copyVal := utils.CopyString(c.Params("message"))
		c.Locals("msg","1:"+copyVal )
		return c.Next()
	}, func(c *fiber.Ctx) error {
		result := fmt.Sprintf("2:%s", c.Locals("msg"))
		return c.SendString(result)
	})
	app.Listen(":3000")
}
