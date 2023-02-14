package domain

import (
	"errors"
	"time"

	"github.com/approvers/qip/pkg/utils/id"
)

type Post struct {
	ID         id.SnowFlakeID
	Body       string
	Type       string
	UserID     id.SnowFlakeID // 作成者のID
	MergeCount int            // マージ数のカウント
	Reactions  []interface{}
	Visibility string // 公開範囲
	CreatedAt  time.Time
}

func NewPost(
	ID id.SnowFlakeID,
	Body string,
	Type string,
	UserID id.SnowFlakeID,
	ReNoteCount int,
	Reactions []interface{},
	Visibility string,
	CreatedAt time.Time,
) (*Post, error) {
	if len(Body) > 3000 {
		return nil, errors.New("body too long")
	}
	return &Post{
		ID:         ID,
		Body:       Body,
		Type:       Type,
		UserID:     UserID,
		MergeCount: ReNoteCount,
		Reactions:  Reactions,
		Visibility: Visibility,
		CreatedAt:  CreatedAt,
	}, nil
}

func (n *Post) ReNoteCountUp() {
	n.MergeCount++
}

func (n *Post) ReNoteCountDown() {
	n.MergeCount--
}
