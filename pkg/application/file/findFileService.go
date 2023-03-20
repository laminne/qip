package file

import (
	"fmt"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/errorType"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type IFindFileService interface {
	FindByID(id id.SnowFlakeID) (*FileData, error)
	FindByUploaderID(id id.SnowFlakeID) ([]FileData, error)
}

type FindFileService struct {
	fileRepository repository.FileRepository
}

func NewFindFileService(r repository.FileRepository) *FindFileService {
	return &FindFileService{fileRepository: r}
}

func (f FindFileService) convert(fi []domain.File) []FileData {
	data := make([]FileData, len(fi))
	for i, v := range fi {
		data[i] = *NewFileData(v)
	}

	return data
}

func (f FindFileService) FindByID(id id.SnowFlakeID) (*FileData, error) {
	file, err := f.fileRepository.FindByID(id)
	fmt.Printf("%#+v", file)
	if err != nil {
		return nil, errorType.NewErrNotFound("FindFileService", "file not found")
	}

	return NewFileData(*file), nil
}

func (f FindFileService) FindByUploaderID(id id.SnowFlakeID) ([]FileData, error) {
	files, err := f.fileRepository.FindByUploaderID(id)
	if err != nil {
		return []FileData{}, errorType.NewErrNotFound("FindFileService", "file not found")
	}

	return f.convert(files), nil
}
