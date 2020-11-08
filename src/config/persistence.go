package config

import (
	"crud-without-database-golang/src/helpers"
	"fmt"
)

// Persistence ...
type Persistence struct {
	Name string
}

// NewPersistence ...
func (app *AppConfig) NewPersistence(name string) *Persistence {
	return &Persistence{
		Name: fmt.Sprintf("%s.json", name),
	}
}

// InitPersistence ...
func (persist *Persistence) InitPersistence() {
	helpers.CreateFile(persist.Name)
}
