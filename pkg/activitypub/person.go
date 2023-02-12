package activitypub

import (
	"encoding/json"
	"fmt"
	"github.com/laminne/notepod/pkg/types"
)

func Person(args types.PersonResponseArgs) []byte {
	res := types.PersonResponseJSONLD{
		Context: []interface{}{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		Type:        "Person",
		ID:          fmt.Sprintf("https://%s/users/%s", InstanceFQDN, args.ID),
		Inbox:       fmt.Sprintf("https://%s/inbox", InstanceFQDN),
		Outbox:      fmt.Sprintf("https://%s/users/%s/outbox", InstanceFQDN, args.ID),
		Followers:   fmt.Sprintf("https://%s/users/%s/followers", InstanceFQDN, args.ID),
		Following:   fmt.Sprintf("https://%s/users/%s/following", InstanceFQDN, args.ID),
		Featured:    fmt.Sprintf("https://%s/users/%s/collections/featured", InstanceFQDN, args.ID),
		SharedInbox: fmt.Sprintf("https://%s/inbox", InstanceFQDN),
		Endpoints: struct {
			SharedInbox string `json:"sharedInbox"`
		}{
			SharedInbox: fmt.Sprintf("https://%s/inbox", InstanceFQDN),
		},
		Url:               fmt.Sprintf("https://%s/@%s", InstanceFQDN, args.UserName),
		PreferredUsername: args.UserName,
		Name:              args.UserScreenName,
		Summary:           args.Summary,
		Icon: struct {
			Type      string      `json:"type"`
			Url       string      `json:"url"`
			Sensitive bool        `json:"sensitive"`
			Name      interface{} `json:"name"`
		}{
			Type:      "Image",
			Url:       args.Icon.Url,
			Sensitive: args.Icon.Sensitive,
			Name:      nil,
		},
		Image: struct {
			Type      string      `json:"type"`
			Url       string      `json:"url"`
			Sensitive bool        `json:"sensitive"`
			Name      interface{} `json:"name"`
		}{
			Type:      "Image",
			Url:       args.Image.Url,
			Sensitive: args.Image.Sensitive,
			Name:      nil,
		},
		Tag: []struct {
			Type string `json:"type"`
			Href string `json:"href"`
			Name string `json:"name"`
		}(args.Tag),
		ManuallyApprovesFollowers: args.ManuallyApprovesFollowers,
		Discoverable:              true,
		PublicKey: struct {
			ID           string `json:"id"`
			Type         string `json:"type"`
			Owner        string `json:"owner"`
			PublicKeyPem string `json:"publicKeyPem"`
		}{
			ID:           fmt.Sprintf("https://%s/users/%s#main-key", InstanceFQDN, args.ID),
			Type:         "Key",
			Owner:        fmt.Sprintf("https://%s/users/%s", InstanceFQDN, args.ID),
			PublicKeyPem: args.PublicKey,
		},
		IsCat:        false,
		VcardBday:    "",
		VcardAddress: "",
	}

	var context = types.PersonResponseContext{
		ManuallyApprovesFollowers: "as:manuallyApprovesFollowers",
		Sensitive:                 "as:sensitive",
		Hashtag:                   "as:Hashtag",
		QuoteUrl:                  "as:quoteUrl",
		Toot:                      "http://joinmastodon.org/ns#",
		Emoji:                     "toot:Emoji",
		Featured:                  "toot:featured",
		Discoverable:              "toot:discoverable",
		Schema:                    "http://schema.org#",
		PropertyValue:             "schema:PropertyValue",
		Value:                     "schema:value",
		Misskey:                   "https://misskey-hub.net/ns#",
		MisskeyContent:            "misskey:_misskey_content",
		MisskeyQuote:              "misskey:_misskey_quote",
		MisskeyReaction:           "misskey:_misskey_reaction",
		MisskeyVotes:              "misskey:_misskey_votes",
		MisskeyTalk:               "misskey:_misskey_talk",
		IsCat:                     "misskey:isCat",
		Vcard:                     "http://www.w3.org/2006/vcard/ns#",
	}

	res.Context = append(res.Context, context)

	j, _ := json.Marshal(res)

	return j
}
