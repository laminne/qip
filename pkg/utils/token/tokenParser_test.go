package token

import (
	"fmt"
	"testing"

	"github.com/approvers/qip/pkg/utils/id"

	"github.com/stretchr/testify/assert"
)

func TestJWTTokenParser_Parse(t *testing.T) {
	key := secureRandom(32)
	g := NewJWTTokenGenerator(key)
	p := NewJWTTokenParser(key)

	token, err := g.NewToken("112233")
	if err != nil {
		t.Fail()
	}
	fmt.Println(token)

	parsed, err := p.Parse(token)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, id.SnowFlakeID("112233"), parsed)

}
