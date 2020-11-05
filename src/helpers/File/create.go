package file

import "os"

func Create(filename string) os.File {
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	return *file
}
