package main

import (
	"fmt"
	"log"

	"eclecticlly/web-stack/bootstrap"
	"eclecticlly/web-stack/pkg/env"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "8000"))))
}
