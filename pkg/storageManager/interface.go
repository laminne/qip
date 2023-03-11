package storageManager

import "io"

// IStorageManager ファイルの保存/取得/管理を行う
type IStorageManager interface {
	Create(fileName string, file io.Reader) (string, error)
}
