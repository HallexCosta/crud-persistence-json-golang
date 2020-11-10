package main

import (
	"crud-without-database-golang/src/feature"
)

// Routes ...
func Routes() {
	app.Get("/users", feature.ListUsers)
	app.Get("/users/:id", feature.ListUser)
	app.Post("/users", feature.CreateUser)
	app.Put("/users/:id", feature.EditUser)
	app.Delete("/users/:id", feature.DeleteUser)
}
