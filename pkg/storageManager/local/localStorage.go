package local

import (
	"fmt"
	"io"
	"os"
)

type StorageManager struct {
}

func NewStorageManager() *StorageManager {
	return &StorageManager{}
}

func (s StorageManager) Create(path string, fileName string, file io.Reader) error {
	openFile, err := os.OpenFile(fmt.Sprintf("%s%s", path, fileName), os.O_CREATE|os.O_RDWR, 0660)
	defer func() {
		_ = openFile.Close()
	}()
	if err != nil {
		return err
	}

	_, err = io.Copy(openFile, file)
	if err != nil {
		return err
	}

	return nil
}
