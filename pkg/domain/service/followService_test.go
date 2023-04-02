package service

import (
	"testing"
	"time"

	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/repository/dummy"
	"github.com/stretchr/testify/assert"
)

func TestFollowService_Exists(t *testing.T) {
	d, _ := domain.NewFollow("123", "456", time.Now())
	a := make([]domain.Follow, 1)
	a[0] = *d
	repo := dummy.NewUserRepository(*new([]domain.User), a)

	followService := NewFollowService(repo)

	// 存在しないとき
	not, _ := domain.NewFollow("987", "654", time.Now())
	assert.Equal(t, false, followService.Exists(*not))

	// 存在するとき
	assert.Equal(t, true, followService.Exists(*d))
}
