package instance

import (
	"time"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type State = int

const (
	NormalInstance = iota
	BlockedInstance
)

type InstanceData struct {
	id              id.SnowFlakeID
	name            string
	softwareName    string
	softwareVersion string
	host            string
	description     string
	state           int
	createdAt       time.Time
	updateAt        *time.Time
}

func NewInstanceData(i domain.Instance) *InstanceData {
	return &InstanceData{
		id:              i.GetID(),
		name:            i.GetName(),
		softwareName:    i.GetSoftwareName(),
		softwareVersion: i.GetSoftwareVersion(),
		host:            i.GetHost(),
		description:     i.GetDescription(),
		state: func() int {
			if i.IsInstanceBlocked() {
				return 1
			}
			return 0
		}(),
		createdAt: i.GetCreatedAt(),
		updateAt:  i.GetUpdatedAt(),
	}
}

func (i InstanceData) Id() id.SnowFlakeID {
	return i.id
}

func (i InstanceData) Name() string {
	return i.name
}

func (i InstanceData) SoftwareName() string {
	return i.softwareName
}

func (i InstanceData) SoftwareVersion() string {
	return i.softwareVersion
}

func (i InstanceData) Host() string {
	return i.host
}

func (i InstanceData) Description() string {
	return i.description
}

func (i InstanceData) State() int {
	return i.state
}

func (i InstanceData) CreatedAt() time.Time {
	return i.createdAt
}

func (i InstanceData) UpdateAt() *time.Time {
	return i.updateAt
}
