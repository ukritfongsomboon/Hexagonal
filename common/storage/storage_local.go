package storage

import (
	"os"
	"path/filepath"
)

// # When Local disk storage is not use
type storageApp struct {
	disk string
}

func NewAppStorage() AppStorage {
	return storageApp{}
}

func (sr storageApp) Write(filename string, data []byte) error {
	// # Split filename and directory
	dir, _ := filepath.Split(filename)

	// # When directory is not root path of application
	if dir != "" {
		// # Create Recurcive directory
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// # Create File
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	// # Finaly To Close File
	defer f.Close()

	// # Write File
	_, err2 := f.Write(data)

	if err2 != nil {
		return err
	}
	// # Succes to write file
	return nil
}

func (sr storageApp) Read(filename string) error {
	return nil
}

func (sr storageApp) Delete(filename string) error {
	return nil
}
