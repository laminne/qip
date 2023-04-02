package activitypub

import (
	"fmt"

	"github.com/laminne/qip/pkg/activitypub/types"
)

func Person(args types.PersonResponseArgs) types.PersonResponseJSONLD {
	res := types.PersonResponseJSONLD{
		Context: []interface{}{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		Type:        "Person",
		ID:          fmt.Sprintf("https://%s/users/%s", args.FQDN, args.ID),
		Inbox:       fmt.Sprintf("https://%s/users/%s/inbox", args.FQDN, args.ID),
		Outbox:      fmt.Sprintf("https://%s/users/%s/outbox", args.FQDN, args.ID),
		Followers:   fmt.Sprintf("https://%s/users/%s/followers", args.FQDN, args.ID),
		Following:   fmt.Sprintf("https://%s/users/%s/following", args.FQDN, args.ID),
		Featured:    fmt.Sprintf("https://%s/users/%s/collections/featured", args.FQDN, args.ID),
		SharedInbox: fmt.Sprintf("https://%s/inbox", args.FQDN),
		Endpoints: struct {
			SharedInbox string `json:"sharedInbox"`
		}{
			SharedInbox: fmt.Sprintf("https://%s/inbox", args.FQDN),
		},
		Url:               fmt.Sprintf("https://%s/@%s", args.FQDN, args.UserName),
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
			ID:           fmt.Sprintf("https://%s/users/%s#main-key", args.FQDN, args.ID),
			Type:         "Key",
			Owner:        fmt.Sprintf("https://%s/users/%s", args.FQDN, args.ID),
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
		Vcard:                     "http://www.w3.org/2006/vcard/ns#",
	}

	res.Context = append(res.Context, context)
	return res
}
