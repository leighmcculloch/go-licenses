// Package licenses contains a list of open source licenses and their text.
package licenses // import "4d63.com/licenses"

func Count() int {
	return len(fileNames)
}

func Names() []string {
	return fileNames
}

func Get(name string) ([]byte, bool) {
	bytes, ok := files[name]
	return bytes, ok
}
