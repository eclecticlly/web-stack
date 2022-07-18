package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// APIRouter for REST API
type APIRouter struct {
}

// InstallRouter export outer
func (h APIRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})
}

// NewAPIRouter initializer
func NewAPIRouter() *APIRouter {
	return &APIRouter{}
}
