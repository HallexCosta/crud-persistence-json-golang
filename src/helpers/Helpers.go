package helpers

import (
	file "crud-without-database-golang/src/helpers/File"
)

// PersistenceFile ...
type PersistenceFile interface {
	CreateFile()
	ReadFile() []byte
	WriteFile(content []byte)
}

// Persistence ...
type Persistence struct {
	Name string
}

// CreateFile ...
func (persist *Persistence) CreateFile() {
	file.Create(persist.Name)
}

// ReadFile ...
func (persist *Persistence) ReadFile() []byte {
	b := file.Read(persist.Name)

	return b
}

// WriteFile ...
func (persist *Persistence) WriteFile(content []byte) {
	file.Write(persist.Name, content)
}
