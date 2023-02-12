package activitypub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcctParser(t *testing.T) {
	host := "example.jp"

	a := AcctParser("@test@example.jp")
	assert.Equal(t, Acct{
		Host:     &(host),
		UserName: "test",
	}, a)

	a = AcctParser("test@example.jp")
	assert.Equal(t, Acct{
		Host:     &(host),
		UserName: "test",
	}, a)
}
