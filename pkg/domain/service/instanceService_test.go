package service

import (
	"testing"
	"time"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/stretchr/testify/assert"
)

func TestInstanceService_Exists(t *testing.T) {
	testData, _ := domain.NewInstance("112233", "example.jp", time.Now())
	testDataArray := make([]domain.Instance, 1)
	testDataArray[0] = *testData
	repo := dummy.NewInstanceRepository(testDataArray)

	instanceService := NewInstanceService(repo)

	// 存在しないとき
	notExist, _ := domain.NewInstance("999", "example.net", time.Now())
	assert.Equal(t, false, instanceService.Exists(*notExist))

	// 存在するとき
	assert.Equal(t, true, instanceService.Exists(*testData))
}
