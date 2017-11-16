package licenses // import "4d63.com/licenses"

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Count() int {
	return len(fileNames)
}

func Names() []string {
	return fileNames
}

func Get(name string) ([]byte, bool) {
	contentsZipped, ok := files[name]
	if !ok {
		return nil, false
	}

	zr, err := gzip.NewReader(bytes.NewBuffer(contentsZipped))
	if err != nil {
		panic(err)
	}

	contents := bytes.Buffer{}
	_, err = io.Copy(&contents, zr)
	if err != nil {
		panic(err)
	}

	zr.Close()

	return contents.Bytes(), true
}
