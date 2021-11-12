package fiber_api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run_server() {
	app := fiber.New(fiber.Config{
		Prefork: true})
	v1 := app.Group("/api/v1")
	Register(v1)

	log.Fatal(app.Listen(":80"))
}
