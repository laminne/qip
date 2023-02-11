package activitypub

import "fmt"

func WebFinger(acct string) string {
	return fmt.Sprintf(`{
		"subject": "acct:%s@np.test.laminne33569.net",
		"links": [
		  {
			"rel": "self",
			"type": "application/activity+json",
			"href": "https://np.test.laminne33569.net/users/%s"
		  },
		  {
			"rel": "http://webfinger.net/rel/profile-page",
			"type": "text/html",
			"href": "https://np.test.laminne33569.net/%s"
		  },
		  {
			"rel": "http://ostatus.org/schema/1.0/subscribe",
			"template": "https://np.test.laminne33569.net/authorize-follow?acct={uri}"
		  }
		]
	  }
	  `, acct, acct, acct)
}
