package feature

import (
	"crud-without-database-golang/src/entity"
	"crud-without-database-golang/src/repository"

	"github.com/gofiber/fiber/v2"
)

// UserRepositoryInterface ...
type UserRepositoryInterface = repository.UserRepositoryInterface

// UserRepository ...
type UserRepository = repository.UserRepository

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
