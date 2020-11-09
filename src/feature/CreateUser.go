package feature

import (
	"crud-without-database-golang/src/entity"

	"github.com/gofiber/fiber/v2"
)

// CreateUser ...
func CreateUser(context *fiber.Ctx) error {
	var user *entity.User
	var userRepository UserRepositoryInterface = new(UserRepository)

	context.BodyParser(user)

	userRepository.Save(user)

	return context.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
