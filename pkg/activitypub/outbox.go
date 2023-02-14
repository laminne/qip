package activitypub

import "github.com/approvers/qip/pkg/activitypub/types"

func UserOutBox() types.UserOutBoxPageJSONLD {
	res := types.UserOutBoxResponseJSONLD{
		Context: []interface{}{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		ID:         "",
		Type:       "",
		TotalItems: 0,
		First:      "",
		Last:       "",
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

	return types.UserOutBoxPageJSONLD{}
}
