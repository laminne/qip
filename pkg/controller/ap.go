package controller

import (
	"github.com/approvers/qip/pkg/activitypub"
	"github.com/approvers/qip/pkg/utils/config"
)

type ActivityPubController struct {
}

func NewActivityPubController() *ActivityPubController {
	return &ActivityPubController{}
}

func (c ActivityPubController) GetNodeInfo() string {
	return activitypub.NodeInfo(config.QipConfig.FQDN)
}

func (c ActivityPubController) GetNodeInfo2() string {
	d := config.QipConfig.Meta
	return activitypub.NodeInfo2(d.Name, d.Description, d.Maintainer.Name, d.Maintainer.Email)
}
