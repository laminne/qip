package file

import (
	"io"
	"time"

	"github.com/approvers/qip/pkg/application/user"

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
	fileService     service.FileService
	findUserService user.FindUserService
	fileRepository  repository.FileRepository
	idGenerator     id.Generator
	storageManager  storageManager.IStorageManager
}

func NewCreateFileService(fileService service.FileService, repository repository.FileRepository, sManager storageManager.IStorageManager, u user.FindUserService) *CreateFileService {
	idGenerator := id.NewSnowFlakeIDGenerator()
	return &CreateFileService{
		fileService:     fileService,
		fileRepository:  repository,
		idGenerator:     idGenerator,
		storageManager:  sManager,
		findUserService: u,
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

	u, err := s.findUserService.FindByID(c.UploaderID)
	if err != nil {
		return nil, err
	}

	if u.IsLocalUser() {
		// ローカルユーザーがアップロードしたファイルは保存する
		p, err := s.storageManager.Create(c.FileName, c.File)
		if err != nil {
			return nil, err
		}
		// ファイルの保存位置(オブジェクトストレージのパス/ファイルへの絶対パスなど)
		f.SetFilePath(p)
	}

	// リモートユーザーのファイルはリンクのみ
	_, err = f.SetFileURL(c.FileURL)
	if err != nil {
		return nil, err
	}

	// ToDo: blurhashを求める
	// DBに保存
	err = s.fileRepository.Create(*f)
	if err != nil {
		return nil, err
	}

	return NewFileData(*f), nil
}
