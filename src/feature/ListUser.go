package feature

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ListUser ...
func ListUser(context *fiber.Ctx) error {
	var userRepository UserRepositoryInterface = new(UserRepository)

	id, _ := strconv.Atoi(context.Params("id"))

	user, isUserFound := userRepository.FindByID(id)

	if isUserFound == false {
		return context.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
