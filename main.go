package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"eclecticlly/web-stack/bootstrap"
	"eclecticlly/web-stack/pkg/env"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static views
var content embed.FS

func main() {
	app := bootstrap.NewApplication()

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(content),
		PathPrefix: "static",
		Browse:     true,
	}))

	go func() {
		log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "8000"))))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	// ...
	fmt.Println("Fiber was successful shutdown.")
}
