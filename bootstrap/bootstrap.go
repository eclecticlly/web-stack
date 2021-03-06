package bootstrap

import (
	"eclecticlly/web-stack/pkg/database"
	"eclecticlly/web-stack/pkg/env"
	"eclecticlly/web-stack/pkg/router"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/jet"
)

// NewApplication initialize application
func NewApplication(content fs.FS) *fiber.App {
	env.SetupEnvFile()
	database.SetupDatabase()
	engine := jet.NewFileSystem(http.FS(content), ".jet")

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			err = ctx.Status(code).SendFile(fmt.Sprintf("views/%d.html", code))
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			return nil
		},
		Views: engine})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())
	router.InstallRouter(app)

	return app
}
