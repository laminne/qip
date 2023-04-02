package activitypub

import (
	"encoding/json"
	"fmt"

	"github.com/laminne/qip/pkg/utils/id"

	"github.com/laminne/qip/pkg/activitypub/types"
)

func WebFinger(acct Acct, fqdn string, userID id.SnowFlakeID) string {
	wf := types.WebFingerResponseJSON{
		Subject: fmt.Sprintf("acct:%s@%s", acct.UserName, *acct.Host),
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
				Href: fmt.Sprintf("https://%s/users/%s", fqdn, userID),
			},
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: fmt.Sprintf("https://%s/@%s", fqdn, acct.UserName),
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: fmt.Sprintf("https://%s/authorize-follow?acct={uri}", fqdn),
			},
		}),
	}

	j, _ := json.Marshal(wf)

	return string(j)
}
