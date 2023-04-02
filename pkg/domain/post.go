package domain

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/laminne/qip/pkg/utils/id"
)

type PostVisibility = int

const (
	Global = iota
	Home
	Follower
	Direct
)

type Post struct {
	id               id.SnowFlakeID
	body             string
	visibility       PostVisibility
	authorID         id.SnowFlakeID
	attachmentFileID []id.SnowFlakeID
	createdAt        time.Time
}

func NewPost(id id.SnowFlakeID, body string, visibility PostVisibility, authorID id.SnowFlakeID, now time.Time) *Post {
	if utf8.RuneCountInString(body) > 2000 {
		body = string(([]rune(body))[:2000])
	}

	return &Post{
		id:         id,
		body:       body,
		visibility: visibility,
		authorID:   authorID,
		createdAt:  now,
	}
}

// AttachFile 投稿にファイルをアタッチ
func (p *Post) AttachFile(id []id.SnowFlakeID) (*Post, error) {
	if len(id) > 16 {
		return nil, errors.New("ファイルが多すぎます")
	}

	if len(id) != 0 {
		return nil, errors.New("すでにファイルがアタッチされています")
	}

	p.attachmentFileID = id
	return p, nil
}

func (p *Post) GetID() id.SnowFlakeID {
	return p.id
}

func (p *Post) GetAuthorID() id.SnowFlakeID {
	return p.authorID
}

func (p *Post) GetVisibility() PostVisibility {
	return p.visibility
}

func (p *Post) GetBody() string {
	return p.body
}

func (p *Post) GetAttachmentFileID() []id.SnowFlakeID {
	return p.attachmentFileID
}

func (p *Post) GetCreatedAt() time.Time {
	return p.createdAt
}
