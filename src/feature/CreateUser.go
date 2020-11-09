package feature

import (
	"crud-without-database-golang/src/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CreateUser ...
func CreateUser(context *fiber.Ctx) error {
	var user *entity.User
	var userRepository UserRepositoryInterface = new(UserRepository)

	context.BodyParser(&user)

	user.ID = int(time.Now().Unix())

	userRepository.Save(user)

	return context.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
