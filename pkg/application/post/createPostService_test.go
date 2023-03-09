package post

import (
	"testing"

	"github.com/approvers/qip/pkg/utils/logger"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostService_Handle(t *testing.T) {
	d := new([]domain.Post)

	repository := dummy.NewPostRepository(*d)
	postService := service.NewPostService(repository)
	createPostService := NewCreatePostService(*postService, repository, logger.NewZapLogger(nil))

	// 成功するとき
	arg := CreatePostCommand{
		Body:       "あいうえお",
		AuthorID:   "112233",
		Visibility: domain.Global,
	}
	_, err := createPostService.Handle(arg)
	assert.Equal(t, nil, err)

	// 失敗するとき
	arg2 := CreatePostCommand{
		Body:       "Hello",
		AuthorID:   "112233",
		Visibility: domain.Global,
	}
	_, err2 := createPostService.Handle(arg2)
	assert.NotEqual(t, nil, err2)
}
