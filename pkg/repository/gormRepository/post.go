package gormRepository

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/entity"
	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/utils/id"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) FindByIDWithUserIcon(id id.SnowFlakeID) (*repository.PostUserFileJoinedData, error) {
	re := entity.PostUserIconJoined{Post: entity.Post{ID: string(id)}}
	r.db.Table("posts").
		Select([]string{"*"}).
		Joins("inner join users on posts.authorid = users.id inner join files on posts.id = files.postid and files.id = users.iconimageid").
		Scan(&re)

	return r.parseJoined(re), nil
}

func (r *PostRepository) FindByAuthorIDWithUserIcon(id id.SnowFlakeID) ([]repository.PostUserFileJoinedData, error) {
	var re []entity.PostUserIconJoined
	r.db.Table("posts").
		Select([]string{"*"}).
		Joins("inner join users on posts.authorid = users.id inner join files on posts.id = files.postid and files.id = users.iconimageid").
		Where("posts.authorid = ?", string(id)).
		Scan(&re)

	res := make([]repository.PostUserFileJoinedData, len(re))
	for j, k := range re {
		res[j] = *r.parseJoined(k)
	}

	return res, nil
}

func NewPostRepository(d *gorm.DB) *PostRepository {
	return &PostRepository{d}
}

func (r *PostRepository) Create(p domain.Post) error {
	e := r.dToE(p)
	res := r.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *PostRepository) FindByID(id id.SnowFlakeID) (*domain.Post, error) {
	p := entity.Post{ID: string(id)}
	res := r.db.Take(&p)
	if res.Error != nil {
		return &domain.Post{}, res.Error
	}
	re := r.eToD(p)
	return &re, nil
}

func (r *PostRepository) FindByAuthorID(id id.SnowFlakeID) ([]domain.Post, error) {
	var p []entity.Post
	res := r.db.Where("authorid = ?", id).Find(&p)
	if res.Error != nil {
		return nil, res.Error
	}

	re := make([]domain.Post, len(p))
	for i, v := range p {
		re[i] = r.eToD(v)
	}
	return re, nil
}

func (r *PostRepository) dToE(p domain.Post) entity.Post {
	return entity.Post{
		ID:         string(p.GetID()),
		Body:       p.GetBody(),
		Visibility: p.GetVisibility(),
		AuthorID:   string(p.GetAuthorID()),
		CreatedAt:  p.GetCreatedAt(),
	}
}

func (r *PostRepository) parseJoined(p entity.PostUserIconJoined) *repository.PostUserFileJoinedData {
	u, _ := domain.NewUser(id.SnowFlakeID(p.User.ID), p.User.Name, id.SnowFlakeID(p.User.InstanceID), p.User.IsLocalUser, p.User.CreatedAt)

	u.SetDisplayName(p.User.DisplayName)
	_, _ = u.SetRole(p.User.Role)
	if p.User.Bio != nil {
		u.SetBio(p.User.Bio)
	}
	if p.User.IsFroze {
		_, _ = u.Freeze()
	}
	_, _ = u.SetInboxURL(p.User.InboxURL)
	_, _ = u.SetOutboxURL(p.User.OutboxURL)
	_, _ = u.SetFollowerURL(p.User.FollowURL)
	_, _ = u.SetFollowerURL(p.User.FollowersURL)
	if p.User.IsLocalUser {
		_, _ = u.SetSecretKey(*p.User.SecretKey)
		_, _ = u.SetPassword(*p.User.Password)
	}
	_, _ = u.SetPublicKey(p.User.PublicKey)
	if p.User.IconImageID != nil {
		u.SetIcon(id.SnowFlakeID(*p.User.IconImageID))
	}
	if p.User.HeaderImageID != nil {
		u.SetHeader(id.SnowFlakeID(*p.User.HeaderImageID))
	}

	f := domain.NewFile(id.SnowFlakeID(p.File.ID), p.File.FileName, id.SnowFlakeID(p.File.UploaderID), p.File.MimeType, p.File.CreatedAt)
	if p.File.PostID != nil {
		f.SetPostID((*id.SnowFlakeID)(p.File.PostID))
	}
	if p.File.FilePath != nil {
		f.SetFilePath(*p.File.FilePath)
	}
	_, _ = f.SetFileURL(p.File.FileURL)
	if p.File.ThumbnailURL != nil {
		_, _ = f.SetThumbnailURL(*p.File.ThumbnailURL)
	}
	f.SetBlurhash(p.File.Blurhash)
	if p.File.IsNSFW {
		_, _ = f.SetNSFW()
	}
	if p.File.UpdatedAt != nil {
		_, _ = f.SetUpdatedAt(*p.File.UpdatedAt)
	}

	post := r.eToD(p.Post)

	return &repository.PostUserFileJoinedData{
		User: u,
		File: f,
		Post: &post,
	}
}

func (r *PostRepository) eToD(e entity.Post) domain.Post {
	// ToDo: AttachmentFileを考慮する
	return *domain.NewPost(id.SnowFlakeID(e.ID), e.Body, e.Visibility, id.SnowFlakeID(e.AuthorID), e.CreatedAt)
}
