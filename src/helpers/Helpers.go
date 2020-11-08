package helpers

import file "crud-without-database-golang/src/helpers/File"

// Persistence ...
type Persistence interface {
	CreateFile(filename string)
	ReadFile(filename string)
	WriteFile(filename string, content string)
}

// CreateFile ...
func CreateFile(filename string) {
	file.Create(filename)
}

// ReadFile ...
func ReadFile(filename string) string {
	b, _ := file.Read(filename)
	return string(b[:])
}

// WriteFile ...
func WriteFile(filename string, content string) {
	file.Write(filename, []byte(content))
}
