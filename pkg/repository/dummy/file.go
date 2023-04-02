package dummy

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/errorType"
	"github.com/laminne/qip/pkg/utils/id"
)

type FileRepository struct {
	data []domain.File
}

func NewFileRepository(data []domain.File) *FileRepository {
	return &FileRepository{data: data}
}

func (r *FileRepository) Create(f domain.File) error {
	for _, v := range r.data {
		if v.GetID() == f.GetID() {
			return errorType.NewErrExists("DummyFileRepository", "File Exists")
		}
	}

	r.data = append(r.data, f)
	return nil
}

func (r *FileRepository) FindByID(id id.SnowFlakeID) (*domain.File, error) {
	for _, v := range r.data {
		if v.GetID() == id {
			return &v, nil
		}
	}

	return nil, errorType.NewErrNotFound("DummyFileRepository", "file not found")
}

func (r *FileRepository) FindByUploaderID(id id.SnowFlakeID) ([]domain.File, error) {
	res := make([]domain.File, 0)
	for _, v := range r.data {
		if v.GetUploaderID() == id {
			res = append(res, v)
		}
	}

	return res, nil
}
