package instance

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository/dummy"
)

func TestCreateInstanceService_Handle(t *testing.T) {
	d := new([]domain.Instance)
	repo := dummy.NewInstanceRepository(*d)
	pService := service.NewInstanceService(repo)
	create := NewCreateInstanceService(*pService, repo)

	// 成功するとき
	arg := CreateInstanceCommand{Host: "example.jp"}
	_, err := create.Handle(arg)
	assert.Equal(t, nil, err)

	// 失敗するとき
	// FIXME: #58で修正

}
