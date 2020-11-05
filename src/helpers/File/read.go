package file

import "io/ioutil"

func Read(filename string) ([]byte, bool) {
	fileContent, _ := ioutil.ReadFile(filename)

	return fileContent, true
}
