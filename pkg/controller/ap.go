package controller

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/laminne/qip/pkg/activitypub/types"
	"github.com/laminne/qip/pkg/utils"

	"github.com/laminne/qip/pkg/application/file"

	"github.com/laminne/qip/pkg/utils/id"

	"github.com/laminne/qip/pkg/application/user"
	"github.com/laminne/qip/pkg/repository"

	"github.com/laminne/qip/pkg/activitypub"
	"github.com/laminne/qip/pkg/utils/config"
)

type ActivityPubController struct {
	findUserService user.FindUserService
	findFileService file.FindFileService
}

func NewActivityPubController(userRepository repository.UserRepository, fileRepository repository.FileRepository) *ActivityPubController {
	return &ActivityPubController{
		findUserService: *user.NewFindUserService(userRepository),
		findFileService: *file.NewFindFileService(fileRepository),
	}
}

func (c ActivityPubController) GetNodeInfo() string {
	return activitypub.NodeInfo(config.QipConfig.FQDN)
}

func (c ActivityPubController) GetNodeInfo2() string {
	d := config.QipConfig.Meta
	return activitypub.NodeInfo2(d.Name, d.Description, d.Maintainer.Name, d.Maintainer.Email)
}

func (c ActivityPubController) GetWebFinger(acct string) (string, error) {
	u := activitypub.AcctParser(acct)
	if *u.Host != config.QipConfig.FQDN {
		return "", errors.New("acct is not local user")
	}

	fmt.Println(*u.Host, u.UserName)
	userData, err := c.findUserService.FindByName(u.UserName)
	if err != nil {
		return "", errors.New("user not found")
	}

	var uID id.SnowFlakeID
	for _, v := range userData {
		if v.IsLocalUser() {
			uID = v.Id()
		}
	}
	if uID == "" {
		return "", errors.New("user not found")
	}

	w := activitypub.WebFinger(u, config.QipConfig.FQDN, uID)
	return w, nil
}

func (c ActivityPubController) GetPerson(userID string) (string, error) {
	u, err := c.findUserService.FindByID(id.SnowFlakeID(userID))
	if err != nil {
		return "", err
	}

	var (
		IconImage   string
		HeaderImage string
		Icon_NSFW   bool
		Header_NSFW bool
	)

	if u.IconImageID() != nil {
		file, err := c.findFileService.FindByID(*u.IconImageID())
		if err != nil {
			return "", err
		}
		IconImage = file.FileURL()
	}
	if u.HeaderImageID() != nil {
		file, err := c.findFileService.FindByID(*u.HeaderImageID())
		if err != nil {
			return "", err
		}
		HeaderImage = file.FileURL()
	}

	arg := types.PersonResponseArgs{
		FQDN:           config.QipConfig.FQDN,
		ID:             string(u.Id()),
		UserName:       u.Name(),
		UserScreenName: u.DisplayName(),
		Summary:        utils.NilFiller[string](u.Bio())[0],
		Icon: struct {
			Url       string
			Sensitive bool
			Name      interface{}
		}{
			IconImage,
			Icon_NSFW,
			"",
		},
		Image: struct {
			Url       string
			Sensitive bool
			Name      interface{}
		}{
			HeaderImage,
			Header_NSFW,
			"",
		},
		Tag:                       nil,
		ManuallyApprovesFollowers: false,
		PublicKey:                 u.PublicKey(),
	}

	res, err := json.Marshal(activitypub.Person(arg))
	if err != nil {
		return "", err
	}
	return string(res), nil
}
