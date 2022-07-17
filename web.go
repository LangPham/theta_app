package main

import (
	"github.com/gofiber/fiber/v2"
)

var Web *ServerWeb

func init() {
	web := ServerWeb{
		name: "webserver",
		app: fiber.New(),
	}
	Web = &web
}

type ServerWeb struct {
	name string
	app *fiber.App
}

func (web *ServerWeb) Start() {
	Sup.wg.Add(1)
	

	web.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	web.app.Listen(":3000")
}

func (web *ServerWeb) Stop() {
	web.app.Shutdown()
	Sup.wg.Done()
}
