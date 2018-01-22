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
func (fm *FileManager) WriteFile(file string, data []byte, permissions int) error {
	perm := int(permissions)
	if perm == 0 {
		perm = int(0644)
	}

	err := ioutil.WriteFile(file, data, os.FileMode(perm))
	if err != nil {
		return errors.Wrap(err, "while writing file")
	}
	return nil
}

// ExistsFile checks if a file exists
func (fm *FileManager) ExistsFile(file string) (bool, error) {
	_, err := ioutil.ReadFile(file)
	if err != nil {
		return false, errors.Wrap(err, "file doesn't exist")
	}
	return true, nil
}
