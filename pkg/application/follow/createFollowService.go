package follow

import (
	"time"

	"github.com/laminne/qip/pkg/errorType"

	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/domain/service"
	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/utils/id"
)

type CreateFollowCommand struct {
	// フォロー元のユーザー
	UserID id.SnowFlakeID
	// フォロー先のユーザー
	TargetID id.SnowFlakeID
}

type ICreateFollowService interface {
	Handle(c CreateFollowCommand) (*FollowData, error)
}

type CreateFollowService struct {
	followService    service.FollowService
	followRepository repository.FollowRepository
}

func NewCreateFollowService(s service.FollowService, repo repository.FollowRepository) *CreateFollowService {
	return &CreateFollowService{
		followService:    s,
		followRepository: repo,
	}
}

func (s *CreateFollowService) Handle(c CreateFollowCommand) (*FollowData, error) {
	f, _ := domain.NewFollow(c.UserID, c.TargetID, time.Now())
	if s.followService.Exists(*f) {
		return nil, errorType.NewErrExists("CreateFollowService", "Already followed")
	}

	err := s.followRepository.Create(*f)
	if err != nil {
		return nil, err
	}

	return NewFollowData(*f), nil
}
