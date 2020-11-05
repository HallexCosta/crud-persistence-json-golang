package main

import (
	"github.com/gofiber/fiber"
)

var app = App()

func App() *fiber.App {
	return fiber.New()
}

func AppListen(p int) {
	app.Listen(p)
}
