package feature

import (
	"crud-without-database-golang/src/entity"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// EditUser ...
func EditUser(context *fiber.Ctx) error {
	var user *entity.User
	var userRepository UserRepositoryInterface = new(UserRepository)

	context.BodyParser(&user)

	id, _ := strconv.Atoi(context.Params("id"))

	_, isUserFound := userRepository.FindByID(id)

	if isUserFound == false {
		return context.Status(400).JSON(fiber.Map{
			"message": "User not found",
			"success": false,
		})
	}

	userRepository.UpdatedByID(id, user)

	return context.Status(200).JSON(fiber.Map{
		"success":      true,
		"user_updated": user,
	})
}
