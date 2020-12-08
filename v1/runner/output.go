package runner

import "os"

func returnFile(filename string) *os.File {
	var OutFile, err = os.Create(filename)
	if err != nil {
		panic(err)
	}
	return OutFile
}
