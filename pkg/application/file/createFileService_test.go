package file

import (
	"bytes"
	"testing"
	"time"

	"github.com/laminne/qip/pkg/application/user"

	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/domain/service"
	"github.com/laminne/qip/pkg/repository/dummy"
	dummyManager "github.com/laminne/qip/pkg/storageManager/dummy"
	"github.com/stretchr/testify/assert"
)

func TestCreateFileService_Handle(t *testing.T) {
	repo := dummy.NewFileRepository(*new([]domain.File))
	u, _ := domain.NewUser("123", "test", "222", true, time.Now())
	users := make([]domain.User, 1)
	users[0] = *u
	uRepo := dummy.NewUserRepository(users, *new([]domain.Follow))
	fileService := *service.NewFileService(repo)
	userService := *user.NewFindUserService(uRepo)
	s := NewCreateFileService(fileService, repo, dummyManager.NewStorageManager("./"), userService)

	// 成功するか
	d := bytes.NewBuffer([]byte("test"))
	_, err := s.Handle(CreateFileCommand{
		FileName:   "test.txt",
		FileURL:    "./",
		UploaderID: "123",
		MimeType:   "text/plain",
		IsNSFW:     false,
		File:       d,
	})
	assert.Equal(t, nil, err)
}
