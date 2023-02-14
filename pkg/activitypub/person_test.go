package activitypub

import (
	"testing"

	"github.com/approvers/qip/pkg/activitypub/types"
	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	p := Person(types.PersonResponseArgs{
		ID:             "1",
		UserName:       "test",
		UserScreenName: "test",
		Summary:        "<p>Hello Fediverse World</p>",
		Icon: struct {
			Url       string
			Sensitive bool
			Name      interface{}
		}{
			Url:       "https://image.example.jp",
			Sensitive: false,
			Name:      nil,
		},
		Image: struct {
			Url       string
			Sensitive bool
			Name      interface{}
		}{
			Url:       "https://image.example.jp",
			Sensitive: false,
			Name:      nil,
		},
		Tag:                       nil,
		ManuallyApprovesFollowers: false,
		PublicKey:                 "-----BEGIN PUBLIC KEY-----\\nDummy\\n-----END PUBLIC KEY-----",
	})

	ex := `{"@context":["https://www.w3.org/ns/activitystreams","https://w3id.org/security/v1",{"manuallyApprovesFollowers":"as:manuallyApprovesFollowers","sensitive":"as:sensitive","Hashtag":"as:Hashtag","quoteUrl":"as:quoteUrl","toot":"http://joinmastodon.org/ns#","Emoji":"toot:Emoji","featured":"toot:featured","discoverable":"toot:discoverable","schema":"http://schema.org#","PropertyValue":"schema:PropertyValue","value":"schema:value","misskey":"https://misskey-hub.net/ns#","_misskey_content":"misskey:_misskey_content","_misskey_quote":"misskey:_misskey_quote","_misskey_reaction":"misskey:_misskey_reaction","_misskey_votes":"misskey:_misskey_votes","_misskey_talk":"misskey:_misskey_talk","isCat":"misskey:isCat","vcard":"http://www.w3.org/2006/vcard/ns#"}],"type":"Person","id":"https://np.test.laminne33569.net/users/1","inbox":"https://np.test.laminne33569.net/inbox","outbox":"https://np.test.laminne33569.net/users/1/outbox","followers":"https://np.test.laminne33569.net/users/1/followers","following":"https://np.test.laminne33569.net/users/1/following","featured":"https://np.test.laminne33569.net/users/1/collections/featured","sharedInbox":"https://np.test.laminne33569.net/inbox","endpoints":{"sharedInbox":"https://np.test.laminne33569.net/inbox"},"url":"https://np.test.laminne33569.net/@test","preferredUsername":"test","name":"test","summary":"\u003cp\u003eHello Fediverse World\u003c/p\u003e","icon":{"type":"Image","url":"https://image.example.jp","sensitive":false,"name":null},"image":{"type":"Image","url":"https://image.example.jp","sensitive":false,"name":null},"tag":null,"manuallyApprovesFollowers":false,"discoverable":true,"publicKey":{"id":"https://np.test.laminne33569.net/users/1#main-key","type":"Key","owner":"https://np.test.laminne33569.net/users/1","publicKeyPem":"-----BEGIN PUBLIC KEY-----\\nDummy\\n-----END PUBLIC KEY-----"},"isCat":false,"vcard:bday":"","vcard:Address":""}`

	assert.Equal(t, []byte(ex), p)

}
