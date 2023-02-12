package id

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSnowFlakeIDGenerator(t *testing.T) {
	expect := SnowFlakeID("6733130380397772800")
	ti, _ := time.Parse(time.RFC3339, "2020-12-02T15:04:05+09:00")

	generator := Generator(&SnowFlakeIDGenerator{})
	assert.Equal(t, expect, generator.NewID(ti))
}
