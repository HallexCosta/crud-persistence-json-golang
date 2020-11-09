package repository

import (
	"crud-without-database-golang/src/config"
	"crud-without-database-golang/src/entity"
	"crud-without-database-golang/src/helpers"
	"encoding/json"
	"fmt"
)

// PersistenceFile ...
type PersistenceFile = helpers.PersistenceFile

// Persistence ...
type Persistence = helpers.Persistence

// UserRepository ...
type UserRepository struct {
}

// Save ...
func (userRepository *UserRepository) Save(user *entity.User) {
	var users []*entity.User

	var persist PersistenceFile = &Persistence{
		Name: config.ImportAppConfig().Persistence.Name,
	}

	usersJSON := persist.ReadFile()

	_ = json.Unmarshal([]byte(usersJSON), &users)

	users = append(users, user)

	usersAddedNewUserJSON, _ := json.Marshal(users)

	persist.WriteFile(usersAddedNewUserJSON)
}

// FindAll ...
func (userRepository *UserRepository) FindAll() []*entity.User {
	var users []*entity.User

	var persist PersistenceFile = &Persistence{
		Name: config.ImportAppConfig().Persistence.Name,
	}

	_ = json.Unmarshal([]byte(persist.ReadFile()), &users)

	fmt.Printf("%+v", users)

	return users
}
