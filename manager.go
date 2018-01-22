package fm

// Manager interface that defines a file manager
type Manager interface {
	OpenFile(file string) (string, error)
	WriteFile(file string, data []byte, permissions int) error
	ExistsFile(file string) (bool, error)
}
