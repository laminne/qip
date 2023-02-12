package id

import (
	"strconv"
	"time"
)

type SnowFlakeID string

type Generator interface {
	NewID(t time.Time) SnowFlakeID
}

func NewSnowFlakeIDGenerator() *SnowFlakeIDGenerator {
	return &SnowFlakeIDGenerator{}
}

type SnowFlakeIDGenerator struct {
}

// SnowFlakeEpoch Epoch: 2020/04/01 00:00:00 (UTC+9/JST)
const SnowFlakeEpoch = 1585666800

func (i *SnowFlakeIDGenerator) NewID(t time.Time) SnowFlakeID {
	var snowFlakeID int64
	var unixDate int64 = t.UnixMilli()
	var date int64 = unixDate - SnowFlakeEpoch
	var workerID int64 = 0
	var processID int64 = 0
	var increment int64 = 0

	snowFlakeID = (((date << 22) + (workerID << 17)) + (processID << 12)) + (increment << 0)

	return SnowFlakeID(strconv.Itoa(int(snowFlakeID)))
}
