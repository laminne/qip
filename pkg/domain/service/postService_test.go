package service

import (
	"testing"
	"time"

	"github.com/laminne/qip/pkg/repository/dummy"
	"github.com/stretchr/testify/assert"

	"github.com/laminne/qip/pkg/domain"
)

func TestPostService_Exists(t *testing.T) {
	testData := *domain.NewPost("112233", "hello hello", domain.Home, "332211", time.Now())
	testDataArray := make([]domain.Post, 1)
	testDataArray[0] = testData
	repo := dummy.NewPostRepository(testDataArray)

	postService := NewPostService(repo)

	// 存在しないとき
	notExist := domain.NewPost("999999", "helloworld", domain.Home, "332211", time.Now())
	assert.Equal(t, false, postService.Exists(notExist))

	// 存在するとき
	exist := testData
	assert.Equal(t, true, postService.Exists(&exist))
}
