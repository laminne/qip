package domain

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/laminne/qip/pkg/utils/id"
)

type InstanceState = int

const (
	NormalInstanceState = iota
	BlockedInstance
)

type Instance struct {
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

func NewInstance(id id.SnowFlakeID, host string, now time.Time) (*Instance, error) {
	if len(host) > 128 {
		return nil, errors.New("ホスト名が長すぎます")
	}
	return &Instance{
		id:        id,
		host:      host,
		createdAt: now,
	}, nil
}

func (i *Instance) SetName(name string) (*Instance, error) {
	if utf8.RuneCountInString(name) > 64 || utf8.RuneCountInString(name) <= 0 {
		name = string(([]rune(name))[:64])
	}
	i.name = name
	return i, nil
}

func (i *Instance) SetSoftwareName(name string) (*Instance, error) {
	if utf8.RuneCountInString(name) > 64 || utf8.RuneCountInString(name) <= 0 {
		name = string(([]rune(name))[:64])
	}
	i.softwareName = name
	return i, nil
}

func (i *Instance) SetSoftwareVersion(version string) (*Instance, error) {
	if utf8.RuneCountInString(version) > 64 || utf8.RuneCountInString(version) <= 0 {
		version = string(([]rune(version))[:64])
	}
	i.softwareVersion = version
	return i, nil
}

func (i *Instance) SetDescription(description string) (*Instance, error) {
	if utf8.RuneCountInString(description) > 2000 || utf8.RuneCountInString(description) <= 0 {
		description = string(([]rune(description))[:2000])
	}
	i.description = description
	return i, nil
}

func (i *Instance) InstanceBlock() (*Instance, error) {
	if i.state == BlockedInstance {
		return nil, errors.New("すでにブロックされています")
	}

	i.state = NormalInstanceState
	return i, nil
}

func (i *Instance) RemoveInstanceBlock() (*Instance, error) {
	if i.state == NormalInstanceState {
		return nil, errors.New("すでにブロック解除されています")
	}

	i.state = BlockedInstance
	return i, nil
}

func (i *Instance) SetUpdatedAt(at time.Time) (*Instance, error) {
	i.updateAt = &at
	return i, nil
}

func (i *Instance) GetID() id.SnowFlakeID {
	return i.id
}

func (i *Instance) GetName() string {
	return i.name
}

func (i *Instance) GetSoftwareName() string {
	return i.softwareName
}

func (i *Instance) GetSoftwareVersion() string {
	return i.softwareVersion
}

func (i *Instance) GetHost() string {
	return i.host
}

func (i *Instance) GetDescription() string {
	return i.description
}

func (i *Instance) IsInstanceBlocked() bool {
	return i.state == BlockedInstance
}

func (i *Instance) GetCreatedAt() time.Time {
	return i.createdAt
}

func (i *Instance) GetUpdatedAt() *time.Time {
	return i.updateAt
}
