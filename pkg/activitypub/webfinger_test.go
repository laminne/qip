package activitypub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebFinger(t *testing.T) {
	a := AcctParser("testuser@np.test.laminne33569.net")
	wf := WebFinger(a, "np.test.laminne33569.net", "123")

	ex := `{"subject":"acct:testuser@np.test.laminne33569.net","links":[{"rel":"self","type":"application/activity+json","href":"https://np.test.laminne33569.net/users/123"},{"rel":"http://webfinger.net/rel/profile-page","type":"text/html","href":"https://np.test.laminne33569.net/@testuser"},{"rel":"http://ostatus.org/schema/1.0/subscribe","template":"https://np.test.laminne33569.net/authorize-follow?acct={uri}"}]}`
	assert.Equal(t, ex, wf)
}
