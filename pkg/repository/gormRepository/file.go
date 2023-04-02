package gormRepository

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/entity"
	"github.com/laminne/qip/pkg/utils/id"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) Create(f domain.File) error {
	e := r.dToE(f)
	res := r.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *FileRepository) FindByID(id id.SnowFlakeID) (*domain.File, error) {
	f := entity.File{ID: string(id)}
	res := r.db.Take(&f)
	if res.Error != nil {
		return nil, res.Error
	}
	re := r.eToD(f)
	return &re, nil
}

func (r *FileRepository) FindByUploaderID(id id.SnowFlakeID) ([]domain.File, error) {
	q := entity.File{UploaderID: string(id)}
	var f []entity.File
	res := r.db.Where(q).Find(&f)
	if res.Error != nil {
		return []domain.File{}, res.Error
	}

	re := make([]domain.File, len(f))
	for i, v := range f {
		re[i] = r.eToD(v)
	}
	return re, nil
}

func (r *FileRepository) dToE(d domain.File) entity.File {
	return entity.File{
		ID:           string(d.GetID()),
		FileName:     d.GetFileName(),
		FilePath:     d.GetFilePath(),
		FileURL:      d.GetFileURL(),
		ThumbnailURL: d.GetThumbnailURL(),
		Blurhash:     d.GetBlurhash(),
		IsNSFW:       d.IsNSFW(),
		MimeType:     d.GetMimeType(),
		CreatedAt:    d.GetCreatedAt(),
		UpdatedAt:    d.GetUpdatedAt(),
		UploaderID:   string(d.GetUploaderID()),
		PostID: func() *string {
			if d.GetPostID() != nil {
				return (*string)(d.GetPostID())
			}
			return (*string)(nil)
		}(),
	}
}

func (r *FileRepository) eToD(e entity.File) domain.File {
	f := domain.NewFile(id.SnowFlakeID(e.ID), e.FileName, id.SnowFlakeID(e.UploaderID), e.MimeType, e.CreatedAt)
	if e.PostID != nil {
		f.SetPostID((*id.SnowFlakeID)(e.PostID))
	}
	if e.FilePath != nil {
		f.SetFilePath(*e.FilePath)
	}
	_, _ = f.SetFileURL(e.FileURL)
	if e.ThumbnailURL != nil {
		_, _ = f.SetThumbnailURL(*e.ThumbnailURL)
	}
	f.SetBlurhash(e.Blurhash)
	if e.IsNSFW {
		_, _ = f.SetNSFW()
	}
	if e.UpdatedAt != nil {
		_, _ = f.SetUpdatedAt(*e.UpdatedAt)
	}
	return *f
}
