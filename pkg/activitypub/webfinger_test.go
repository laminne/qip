package activitypub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebFinger(t *testing.T) {
	wf, _ := WebFinger("testuser@np.test.laminne33569.net", "np.test.laminne33569.net")

	ex := `{"subject":"acct:testuser@np.test.laminne33569.net","links":[{"rel":"self","type":"application/activity+json","href":"https://np.test.laminne33569.net/users/testuser"},{"rel":"http://webfinger.net/rel/profile-page","type":"text/html","href":"https://np.test.laminne33569.net/@testuser"},{"rel":"http://ostatus.org/schema/1.0/subscribe","template":"https://np.test.laminne33569.net/authorize-follow?acct={uri}"}]}`
	assert.Equal(t, ex, wf)
}
