package extract

import (
	"bufio"
	"io"
	"log"
	"os"
)

type flags struct {
	isBodyFound bool
}

func (f *flags) reset() {
	f.isBodyFound = false
}

var f = &flags{}

func Class(file *os.File) map[string]struct{} {
	classesMap := make(map[string]struct{})
	reader := bufio.NewReader(file)
	f.reset()

	for {
		_, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	return classesMap
}
