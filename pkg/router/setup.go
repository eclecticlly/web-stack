package router

import (
	"github.com/gofiber/fiber/v2"
)

// InstallRouter to the app
func InstallRouter(app *fiber.App) {
	setup(app, NewAPIRouter(), NewHTTPRouter())
}
func setup(app *fiber.App, router ...Router) {
	for _, r := range router {
		r.InstallRouter(app)
	}
}
