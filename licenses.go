// Package licenses contains a list of open source licenses and their text.
package licenses // import "4d63.com/licenses"

func Count() int {
	return len(fileNames)
}

func Names() []string {
	return fileNames
}

func Text(licenseName string) ([]byte, bool) {
	bytes, ok := files[licenseName]
	return bytes, ok
}
