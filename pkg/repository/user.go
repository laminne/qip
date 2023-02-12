package repository

import (
	"github.com/laminne/notepod/pkg/models"
	"github.com/laminne/notepod/pkg/utils/id"
)

type UserRepository interface {
	Create(u models.User) (models.User, error)
	Update(u models.User) (models.User, error)
	FindByID(id id.SnowFlakeID) (*models.User, error)
	FindByUserName(n string) (*models.User, error)
	FindByHost(h string) ([]models.User, error)
}
