package file

import (
	"io"
	"time"

	"github.com/approvers/qip/pkg/storageManager"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type CreateFileCommand struct {
	FileName   string
	FileURL    string
	UploaderID id.SnowFlakeID
	MimeType   string
	IsNSFW     bool
	File       io.Reader
}

type ICreateFileService interface {
	Handle(c CreateFileCommand) (*FileData, error)
}

type CreateFileService struct {
	fileService    service.FileService
	fileRepository repository.FileRepository
	idGenerator    id.Generator
	storageManager storageManager.IStorageManager
}

func NewCreateFileService(fileService service.FileService, repository repository.FileRepository, sManager storageManager.IStorageManager) *CreateFileService {
	idGenerator := id.NewSnowFlakeIDGenerator()
	return &CreateFileService{
		fileService:    fileService,
		fileRepository: repository,
		idGenerator:    idGenerator,
		storageManager: sManager,
	}
}

func (s *CreateFileService) Handle(c CreateFileCommand) (*FileData, error) {
	now := time.Now()

	fID := s.idGenerator.NewID(now)
	f := domain.NewFile(fID, c.FileName, c.UploaderID, c.MimeType, now)
	if c.IsNSFW {
		_, err := f.SetNSFW()
		if err != nil {
			return nil, err
		}
	}
	_, err := f.SetFileURL(c.FileURL)
	if err != nil {
		return nil, err
	}
	// ToDo: blurhashを求める

	// DBに保存
	err = s.fileRepository.Create(*f)
	if err != nil {
		return nil, err
	}

	// ストレージに保存する
	// ToDo: リモートの画像の考慮
	/*
		FIXME: URLをパースする処理があったほうがいいかもしれない
			例: file:///home/laminne/test.png -> /home/laminne/test.png に保存
				https://example.jp/test.png -> そのまま
	*/
	err = s.storageManager.Create(c.FileURL, c.FileName, c.File)
	if err != nil {
		return nil, err
	}

	return NewFileData(*f), nil
}
