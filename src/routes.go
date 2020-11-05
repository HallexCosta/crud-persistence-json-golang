package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	app.Get("/", func(context *fiber.Ctx) error {
		User := new(struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Age      int    `json:"age"`
		})

		context.BodyParser(User)

		// Destructuring
		Name, Email, Password, Age := User.Name, User.Email, User.Password, User.Age
		fmt.Printf("User {\n  name: %s\n  email: %s\n  password: %s\n  age: %x\n}\n", Name, Email, Password, Age)

		return context.JSON(fiber.Map{
			"data":    User,
			"success": true,
			"message": "Current Context Success",
		})
	})
}
