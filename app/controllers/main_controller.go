package controllers

import "github.com/gofiber/fiber/v2"

func RenderHello(c *fiber.Ctx) error {

	return c.Render("index.jet", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/main.jet")
}
