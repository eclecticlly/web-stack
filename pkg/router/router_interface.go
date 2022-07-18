package router

import "github.com/gofiber/fiber/v2"

// Router generic type
type Router interface {
	InstallRouter(app *fiber.App)
}
