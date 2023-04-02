package file

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/errorType"
	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/utils/id"
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
