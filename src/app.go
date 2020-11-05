package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var app = App()

func App() *fiber.App {
	return fiber.New()
}

func AppListen(p int) {
	port := fmt.Sprintf(":%d", p)
	app.Listen(port)
}
