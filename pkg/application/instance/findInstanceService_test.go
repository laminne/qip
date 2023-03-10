package instance

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository/dummy"
)

var (
	data                []domain.Instance
	findInstanceService FindInstanceService
)

func init() {
	i, _ := domain.NewInstance("123", "example.jp", time.Now())
	data = append(data, *i)

	findInstanceService = *NewFindInstanceService(dummy.NewInstanceRepository(data))
}

func TestFindInstanceService_FindByID(t *testing.T) {
	// 成功するとき
	a, _ := findInstanceService.FindByID("123")
	assert.Equal(t, NewInstanceData(data[0]), a)

	// 失敗するとき
	_, err := findInstanceService.FindByID("999")
	assert.NotEqual(t, nil, err)
}
