package repository

import (
	"crud-without-database-golang/src/config"
	"crud-without-database-golang/src/entity"
	"crud-without-database-golang/src/helpers"
	"encoding/json"
)

// PersistenceFile ...
type PersistenceFile = helpers.PersistenceFile

// Persistence ...
type Persistence = helpers.Persistence

// UserRepositoryInterface ...
type UserRepositoryInterface interface {
	Save(user *entity.User)
	FindByID(id int) (*entity.User, bool)
	FindAll() []*entity.User
	UpdatedByID(id int, data *entity.User)
	DeleteByID(id int)
}

// UserRepository ...
type UserRepository struct {
}

// Save ...
func (userRepository *UserRepository) Save(user *entity.User) {
	var persist PersistenceFile = &Persistence{
		Name: config.ImportAppConfig().Persistence.Name,
	}

	var users []*entity.User = userRepository.FindAll()

	users = append(users, user)

	updatedUsers, _ := json.MarshalIndent(users, "", "\t")

	persist.WriteFile(updatedUsers)
}

// FindByID ...
func (userRepository *UserRepository) FindByID(id int) (*entity.User, bool) {
	var users []*entity.User = userRepository.FindAll()

	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}

	return &entity.User{}, false
}

// FindAll ...
func (userRepository *UserRepository) FindAll() []*entity.User {
	var persist PersistenceFile = &Persistence{
		Name: config.ImportAppConfig().Persistence.Name,
	}

	var users []*entity.User

	getUsers := persist.ReadFile()

	_ = json.Unmarshal(getUsers, &users)

	return users
}

// UpdatedByID ...
func (userRepository *UserRepository) UpdatedByID(id int, data *entity.User) {
	var persist PersistenceFile = &Persistence{
		Name: config.ImportAppConfig().Persistence.Name,
	}

	var users []*entity.User
	var updatedUser *entity.User

	for _, user := range userRepository.FindAll() {
		if user.ID != id {
			users = append(users, user)
		} else {
			updatedUser = data
			updatedUser.ID = user.ID
		}
	}

	users = append(users, updatedUser)

	usersJSON, _ := json.MarshalIndent(users, "", "\t")

	persist.WriteFile(usersJSON)
}

// DeleteByID ...
func (userRepository *UserRepository) DeleteByID(id int) {
	var persist PersistenceFile = &Persistence{
		Name: config.ImportAppConfig().Persistence.Name,
	}

	var users []*entity.User

	for _, user := range userRepository.FindAll() {
		if user.ID != id {
			users = append(users, user)
		}
	}

	usersJSON, _ := json.MarshalIndent(users, "", "\t")

	persist.WriteFile(usersJSON)
}
