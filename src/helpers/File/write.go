package file

import (
	"io/ioutil"
	"os"
)

func Write(filename string, content []byte) {
	var mode os.FileMode = 0644

	_ = ioutil.WriteFile(filename, content, mode)
}
