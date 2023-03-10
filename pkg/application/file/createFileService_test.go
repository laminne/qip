package file

import (
	"io"
	"testing"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/approvers/qip/pkg/storageManager/local"
	"github.com/stretchr/testify/assert"
)

func TestCreateFileService_Handle(t *testing.T) {
	repo := dummy.NewFileRepository(*new([]domain.File))
	fileService := *service.NewFileService(repo)
	s := NewCreateFileService(fileService, repo, local.NewStorageManager())

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
