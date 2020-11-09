package config

import (
	"crud-without-database-golang/src/helpers"
	"fmt"
)

// PersistenceFile ...
type PersistenceFile = helpers.PersistenceFile

// Persistence ...
type Persistence struct {
	Name string
}

// NewPersistence ...
func (app *AppConfig) NewPersistence(name string) *Persistence {
	app.Persistence.Name = fmt.Sprintf("%s.json", name)
	return &app.Persistence
}

// InitPersistence ...
func (persist *Persistence) InitPersistence() {
	var persistence PersistenceFile = &helpers.Persistence{
		Name: ImportAppConfig().Persistence.Name,
	}

	persistence.CreateFile()
}
