package bootstrap

import (
	"eclecticlly/web-stack/pkg/database"
	"eclecticlly/web-stack/pkg/env"
	"eclecticlly/web-stack/pkg/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/jet"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	database.SetupDatabase()
	engine := jet.New("./views", ".jet")

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
		Views: engine})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())
	router.InstallRouter(app)

	return app
}
