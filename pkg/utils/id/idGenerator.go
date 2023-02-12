package id

import (
	"strconv"
	"time"
)

type SnowFlakeID string

type Generator interface {
	NewID() SnowFlakeID
}

type SnowFlakeIDGenerator struct {
	Time time.Time
}

// SnowFlakeEpoch Epoch: 2020/04/01 00:00:00 (UTC+9/JST)
const SnowFlakeEpoch = 1585666800

func (i *SnowFlakeIDGenerator) NewID() SnowFlakeID {
	var snowFlakeID int64
	var unixDate int64 = i.Time.UnixMilli()
	var date int64 = unixDate - SnowFlakeEpoch
	var workerID int64 = 0
	var processID int64 = 0
	var increment int64 = 0

	snowFlakeID = (((date << 22) + (workerID << 17)) + (processID << 12)) + (increment << 0)

	return SnowFlakeID(strconv.Itoa(int(snowFlakeID)))
}
