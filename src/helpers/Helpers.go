package helpers

import file "crud-without-database-golang/src/helpers/File"

func CreateFile(filename string) {
	file.Create(filename)
}
func ReadFile(filename string) string {
	b, _ := file.Read(filename)
	return string(b[:])
}
func WriteFile(filename string, content string) {
	file.Write(filename, []byte(content))
}
