package types

type PersonResponseArgs struct {
	FQDN           string
	ID             string
	UserName       string
	UserScreenName string
	Summary        string
	Icon           struct {
		Url       string
		Sensitive bool
		Name      interface{} // 基本nullっぽい？
	}
	Image struct {
		Url       string
		Sensitive bool
		Name      interface{}
	}
	Tag []struct {
		Type string
		Href string
		Name string
	}
	ManuallyApprovesFollowers bool
	PublicKey                 string
}

type PersonResponseContext struct {
	ManuallyApprovesFollowers string `json:"manuallyApprovesFollowers"`
	Sensitive                 string `json:"sensitive"`
	Hashtag                   string `json:"Hashtag"`
	QuoteUrl                  string `json:"quoteUrl"`
	Toot                      string `json:"toot"`
	Emoji                     string `json:"Emoji"`
	Featured                  string `json:"featured"`
	Discoverable              string `json:"discoverable"`
	Schema                    string `json:"schema"`
	PropertyValue             string `json:"PropertyValue"`
	Value                     string `json:"value"`
	Misskey                   string `json:"misskey"`
	MisskeyContent            string `json:"_misskey_content"`
	MisskeyQuote              string `json:"_misskey_quote"`
	MisskeyReaction           string `json:"_misskey_reaction"`
	MisskeyVotes              string `json:"_misskey_votes"`
	MisskeyTalk               string `json:"_misskey_talk"`
	IsCat                     string `json:"isCat"`
	Vcard                     string `json:"vcard"`
}

// PersonResponseJSONLD APのユーザー情報(JSON-LD)
type PersonResponseJSONLD struct {
	Context     []interface{} `json:"@context"`    // ごちゃごちゃしてる
	Type        string        `json:"type"`        // Person
	ID          string        `json:"id"`          // /users/:id
	Inbox       string        `json:"inbox"`       // /inbox
	Outbox      string        `json:"outbox"`      // /users/:id/outbox
	Followers   string        `json:"followers"`   // /users/:id/followers
	Following   string        `json:"following"`   // /users/:id/following
	Featured    string        `json:"featured"`    // /users/:id/collections/featured
	SharedInbox string        `json:"sharedInbox"` // /inbox
	Endpoints   struct {
		SharedInbox string `json:"sharedInbox"` // /inbox
	} `json:"endpoints"`
	Url               string `json:"url"`               // /@ユーザー名
	PreferredUsername string `json:"preferredUsername"` // ユーザー名
	Name              string `json:"name"`              // ユーザー表示名
	Summary           string `json:"summary"`           // ユーザーのbio
	Icon              struct {
		Type      string      `json:"type"`
		Url       string      `json:"url"`
		Sensitive bool        `json:"sensitive"`
		Name      interface{} `json:"name"`
	} `json:"icon"`
	Image struct {
		Type      string      `json:"type"`
		Url       string      `json:"url"`
		Sensitive bool        `json:"sensitive"`
		Name      interface{} `json:"name"`
	} `json:"image"`
	Tag []struct {
		Type string `json:"type"`
		Href string `json:"href"`
		Name string `json:"name"`
	} `json:"tag"`
	ManuallyApprovesFollowers bool `json:"manuallyApprovesFollowers"` // 鍵アカウントか
	Discoverable              bool `json:"discoverable"`              // みつけるに表示するか?
	PublicKey                 struct {
		ID           string `json:"id"`    // /users/:id#main-key
		Type         string `json:"type"`  // Key
		Owner        string `json:"owner"` // /users/:id
		PublicKeyPem string `json:"publicKeyPem"`
	} `json:"publicKey"`
	IsCat        bool   `json:"isCat"`
	VcardBday    string `json:"vcard:bday"`
	VcardAddress string `json:"vcard:Address"`
}
