package id

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSnowFlakeIDGenerator(t *testing.T) {
	expect := SnowFlakeID("6733130380397772800")
	time, _ := time.Parse(time.RFC3339, "2020-12-02T15:04:05+09:00")
	generator := Generator(&SnowFlakeIDGenerator{Time: time})
	assert.Equal(t, expect, generator.NewID())
}
