package file

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/repository/dummy"
)

var findFileService FindFileService

func init() {
	f := domain.NewFile("334", "test.jpg", "222", "img/jpeg", time.Now())
	fA := make([]domain.File, 0)
	fA = append(fA, *f)
	r := dummy.NewFileRepository(fA)
	findFileService = *NewFindFileService(r)
}

func TestFindFileService_FindByID(t *testing.T) {
	// 成功するとき
	_, err := findFileService.FindByID("334")
	assert.Equal(t, nil, err)

	// 失敗するとき
	_, err = findFileService.FindByID("111")
	assert.NotEqual(t, nil, err)
}
