package service

import (
	"testing"
	"time"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/stretchr/testify/assert"
)

func TestFileService_Exists(t *testing.T) {
	testData := *domain.NewFile("123", "hello.png", "333", "image/png", time.Now())
	testDataArray := make([]domain.File, 1)
	testDataArray[0] = testData
	repo := dummy.NewFileRepository(testDataArray)

	fileService := NewFileService(repo)

	// 存在しないとき
	notExist := *domain.NewFile("1111111", "ne.jpg", "7777", "image/jpeg", time.Now())
	assert.Equal(t, false, fileService.Exists(notExist))

	// 存在するとき
	exist := testData
	assert.Equal(t, true, fileService.Exists(exist))

}
