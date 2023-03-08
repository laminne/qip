package post

import (
	"time"

	"github.com/approvers/qip/pkg/domain"

	"github.com/approvers/qip/pkg/utils/id"
)

type Visibility = int

const (
	Global = iota
	Home
	Follower
	Direct
)

type PostData struct {
	id               id.SnowFlakeID
	body             string
	visibility       Visibility
	authorID         id.SnowFlakeID
	attachmentFileID []id.SnowFlakeID
	createdAt        time.Time
}

func NewPostData(p domain.Post) *PostData {
	return &PostData{
		id:               p.GetID(),
		body:             p.GetBody(),
		visibility:       p.GetVisibility(),
		authorID:         p.GetAuthorID(),
		attachmentFileID: p.GetAttachmentFileID(),
		createdAt:        p.GetCreatedAt(),
	}
}

func (p PostData) GetID() id.SnowFlakeID {
	return p.id
}

func (p PostData) GetBody() string {
	return p.body
}

func (p PostData) GetVisibility() Visibility {
	return p.visibility
}

func (p PostData) GetAuthorID() id.SnowFlakeID {
	return p.authorID
}

func (p PostData) GetAttachmentFileID() []id.SnowFlakeID {
	return p.attachmentFileID
}

func (p PostData) GetCreatedAt() time.Time {
	return p.createdAt
}
