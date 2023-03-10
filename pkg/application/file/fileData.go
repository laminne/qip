package file

import (
	"github.com/approvers/qip/pkg/domain"
	"time"

	"github.com/approvers/qip/pkg/utils/id"
)

type FileData struct {
	id           id.SnowFlakeID
	uploaderID   id.SnowFlakeID
	postID       *id.SnowFlakeID
	fileName     string
	fileURL      string
	thumbnailURL *string
	blurhash     string
	isNSFW       bool
	mimeType     string
	createdAt    time.Time
	updatedAt    *time.Time
}

func NewFileData(f domain.File) *FileData {
	return &FileData{
		id:           f.GetID(),
		uploaderID:   f.GetUploaderID(),
		postID:       f.GetPostID(),
		fileName:     f.GetFileName(),
		fileURL:      f.GetFileURL(),
		thumbnailURL: f.GetThumbnailURL(),
		blurhash:     f.GetBlurhash(),
		isNSFW:       f.IsNSFW(),
		mimeType:     f.GetMimeType(),
		createdAt:    f.GetCreatedAt(),
		updatedAt:    f.GetUpdatedAt(),
	}
}

func (f FileData) Id() id.SnowFlakeID {
	return f.id
}

func (f FileData) UploaderID() id.SnowFlakeID {
	return f.uploaderID
}

func (f FileData) PostID() *id.SnowFlakeID {
	return f.postID
}

func (f FileData) FileName() string {
	return f.fileName
}

func (f FileData) FileURL() string {
	return f.fileURL
}

func (f FileData) ThumbnailURL() *string {
	return f.thumbnailURL
}

func (f FileData) Blurhash() string {
	return f.blurhash
}

func (f FileData) IsNSFW() bool {
	return f.isNSFW
}

func (f FileData) MimeType() string {
	return f.mimeType
}

func (f FileData) CreatedAt() time.Time {
	return f.createdAt
}

func (f FileData) UpdatedAt() *time.Time {
	return f.updatedAt
}
