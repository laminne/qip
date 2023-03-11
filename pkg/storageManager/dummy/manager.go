package dummy

import (
	"fmt"
	"io"
)

type StorageManager struct {
	basePath string
}

func NewStorageManager(basePath string) *StorageManager {
	return &StorageManager{
		basePath: basePath,
	}
}

func (s StorageManager) Create(fileName string, _ io.Reader) (string, error) {
	return fmt.Sprintf("%s%s", s.basePath, fileName), nil
}
