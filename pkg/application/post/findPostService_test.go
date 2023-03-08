package post

import (
	"testing"
	"time"

	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/stretchr/testify/assert"

	"github.com/approvers/qip/pkg/domain"
)

var (
	data            []domain.Post
	findPostService FindPostService
)

func init() {
	a := domain.NewPost("123", "hello", Global, "111", time.Now())
	b := domain.NewPost("777", "test1", Global, "222", time.Now())
	c := domain.NewPost("888", "world", Global, "222", time.Now())
	data = append(data, *a)
	data = append(data, *b)
	data = append(data, *c)

	repo := dummy.NewPostRepository(data)
	findPostService = *NewFindPostService(repo)
}

func TestFindPostService_FindByID(t *testing.T) {
	// 成功するとき
	a, _ := findPostService.FindByID("123")
	assert.Equal(t, NewPostData(data[0]), a)

	// 失敗するとき
	_, err := findPostService.FindByID("999")
	assert.NotEqual(t, nil, err)
}

func TestFindPostService_FindByAuthorID(t *testing.T) {
	ex := []PostData{*NewPostData(data[1]), *NewPostData(data[2])}

	// 成功するとき
	a, _ := findPostService.FindByAuthorID("222")
	assert.Equal(t, ex, a)

	// 失敗するとき
	b, _ := findPostService.FindByAuthorID("000")
	assert.Equal(t, 0, len(b))

}
