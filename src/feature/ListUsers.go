package feature

import (
	"crud-without-database-golang/src/entity"

	"github.com/gofiber/fiber/v2"
)

// ListUsers ...
func ListUsers(context *fiber.Ctx) error {
	var users []*entity.User

	var userRepository UserRepositoryInterface = new(UserRepository)

	users = userRepository.FindAll()

	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}
