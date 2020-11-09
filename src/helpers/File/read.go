package file

import "io/ioutil"

func Read(filename string) []byte {
	fileContent, _ := ioutil.ReadFile(filename)

	return fileContent
}
