package fs

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/rytsh/liz/file"
)

var API = file.New()

var BasePath = "./"

func AddPath(path string) string {
	return filepath.Join(BasePath, path)
}

func Save(path string, data io.Reader) error {
	return API.SetRawWithReader(AddPath(path), data)
}

func IsExist(path string) bool {
	if _, err := os.Stat(AddPath(path)); errors.Is(err, fs.ErrNotExist) {
		return false
	}

	return true
}

func Delete(path string, force bool) error {
	// delete directory
	if force {
		return os.RemoveAll(AddPath(path))
	}

	// delete file
	return os.Remove(AddPath(path))
}
