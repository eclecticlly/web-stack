package router

import (
	"eclecticlly/web-stack/app/controllers"
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static/*
var content embed.FS

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(content),
		PathPrefix: "static",
		Browse:     true,
	}))

	group := app.Group("", cors.New(), csrf.New())
	group.Get("/", controllers.RenderHello)
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
