package fm

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// FileManager implementation of a Manager
type FileManager struct {
}

// NewFileManager returns a new Manager
func NewFileManager() Manager {
	return &FileManager{}
}

// OpenFile opens the given file and returns the string of it's content
func (fm *FileManager) OpenFile(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", errors.Wrap(err, "While reading file")
	}

	content := string(bytes)
	return content, nil
}

// WriteFile saves the data in the file, with the given permission
// if 0 is sent as permission 0644 will be used to write the file
func (fm *FileManager) WriteFile(file string, data []byte, permissions uint32) error {
	if permissions == 0 {
		permissions = 0644
	}

	if err := ioutil.WriteFile(file, data, os.FileMode(permissions)); err != nil {
		return errors.Wrap(err, "while writing file")
	}
	return nil
}

// ExistsFile checks if a file exists
func (fm *FileManager) ExistsFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}
	return true
}
