package domain

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/laminne/qip/pkg/utils/id"
)

// File ファイル
type File struct {
	id           id.SnowFlakeID
	uploaderID   id.SnowFlakeID
	postID       *id.SnowFlakeID
	fileName     string
	filePath     *string
	fileURL      string
	thumbnailURL *string
	blurhash     string
	isNSFW       bool
	mimeType     string
	createdAt    time.Time
	updatedAt    *time.Time
}

func NewFile(id id.SnowFlakeID, fileName string, uploaderID id.SnowFlakeID, mimeType string, now time.Time) *File {
	fileName = fmt.Sprintf("%s.%s", id, filepath.Ext(fileName))

	return &File{
		id:         id,
		uploaderID: uploaderID,
		fileName:   fileName,
		mimeType:   mimeType,
		isNSFW:     false,
		createdAt:  now,
	}
}

func (f *File) SetFileURL(url string) (*File, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}
	f.fileURL = url
	return f, nil
}

func (f *File) SetThumbnailURL(url string) (*File, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}
	f.thumbnailURL = &url
	return f, nil
}

func (f *File) SetBlurhash(hash string) *File {
	f.blurhash = hash
	return f
}

func (f *File) SetFilePath(p string) *File {
	f.filePath = &p
	return f
}

func (f *File) SetNSFW() (*File, error) {
	if f.isNSFW {
		return nil, errors.New("すでにNSFWに設定されています")
	}

	f.isNSFW = true
	return f, nil
}

func (f *File) UnsetNSFW() (*File, error) {
	if !f.isNSFW {
		return nil, errors.New("すでにNSFWが解除されています")
	}

	f.isNSFW = false
	return f, nil
}

func (f *File) SetMimeType(mime string) (*File, error) {
	f.mimeType = mime
	return f, nil
}

func (f *File) SetUpdatedAt(t time.Time) (*File, error) {
	f.updatedAt = &t
	return f, nil
}

func (f *File) GetID() id.SnowFlakeID {
	return f.id
}

func (f *File) GetUploaderID() id.SnowFlakeID {
	return f.uploaderID
}

func (f *File) GetFileName() string {
	return f.fileName
}

func (f *File) GetFileURL() string {
	return f.fileURL
}

func (f *File) GetFilePath() *string {
	return f.filePath
}

func (f *File) GetThumbnailURL() *string {
	return f.thumbnailURL
}

func (f *File) GetBlurhash() string {
	return f.blurhash
}

func (f *File) IsNSFW() bool {
	return f.isNSFW
}

func (f *File) GetMimeType() string {
	return f.mimeType
}

func (f *File) GetCreatedAt() time.Time {
	return f.createdAt
}

func (f *File) GetUpdatedAt() *time.Time {
	return f.updatedAt
}

func (f *File) GetPostID() *id.SnowFlakeID {
	return f.postID
}

func (f *File) SetPostID(postID *id.SnowFlakeID) {
	f.postID = postID
}
