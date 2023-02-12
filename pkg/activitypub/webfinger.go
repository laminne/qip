package activitypub

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/laminne/notepod/pkg/activitypub/types"
)

func WebFinger(acct string) (string, error) {
	u := AcctParser(acct)
	if *u.Host != InstanceFQDN {
		return "", errors.New("acct is not local user")
	}

	wf := types.WebFingerResponseJSON{
		Subject: fmt.Sprintf("acct:%s", acct),
		Links: []struct {
			Rel      string `json:"rel"`
			Type     string `json:"type,omitempty"`
			Href     string `json:"href,omitempty"`
			Template string `json:"template,omitempty"`
		}([]struct {
			Rel      string
			Type     string
			Href     string
			Template string
		}{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: fmt.Sprintf("https://%s/users/%s", InstanceFQDN, u.UserName),
			},
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: fmt.Sprintf("https://%s/@%s", InstanceFQDN, u.UserName),
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: fmt.Sprintf("https://%s/authorize-follow?acct={uri}", InstanceFQDN),
			},
		}),
	}

	j, _ := json.Marshal(wf)

	return string(j), nil
}
