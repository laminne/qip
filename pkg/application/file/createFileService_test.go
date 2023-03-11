package file

import (
	"io"
	"testing"
	"time"

	"github.com/approvers/qip/pkg/application/user"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/approvers/qip/pkg/storageManager/local"
	"github.com/stretchr/testify/assert"
)

func TestCreateFileService_Handle(t *testing.T) {
	repo := dummy.NewFileRepository(*new([]domain.File))
	u, _ := domain.NewUser("123", "test", "222", true, time.Now())
	users := make([]domain.User, 1)
	users[0] = *u
	uRepo := dummy.NewUserRepository(users)
	fileService := *service.NewFileService(repo)
	userService := *user.NewFindUserService(uRepo)
	s := NewCreateFileService(fileService, repo, local.NewStorageManager(), userService)

	// 成功するか
	d := newDummyReader([]byte("test"))
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

type dummyReader struct {
	data     []byte
	readByte int
}

func newDummyReader(data []byte) *dummyReader {
	return &dummyReader{
		data: data,
	}
}

func (d dummyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		// データ末尾に来たら脱出
		if i+d.readByte == len(d.data) {
			return d.readByte + i, io.EOF
		}
		p[i] = d.data[d.readByte+i]
	}
	return len(p), nil
}
