package activitypub

import (
	"fmt"
)

func NodeInfo(fqdn string) string {
	return fmt.Sprintf(`{
		"links": [
			{
				"rel": "http://nodeinfo.diaspora.software/ns/schema/2.0",
				"href": "https://%v/nodeinfo/2.0"
			}
		]
	}`,
		fqdn)
}

func NodeInfo2(name string, description string, maintainerName string, email string) string {
	return fmt.Sprintf(`{
	"version": "2.0",
	"software": {
	  "name": "Qip",
	  "version": "0.0.1"
	},
	"protocols": [
	  "activitypub"
	],
	"services": {
	  "inbound": [],
	  "outbound": []
	},
	"openRegistrations": false,
	"usage": {
	  "users": {
		"total": 1,
		"activeHalfyear": null,
		"activeMonth": null
	  },
	  "localPosts": 0
	},
	"metadata": {
	  "nodeName": "%s",
	  "nodeDescription": "%s",
	  "maintainer": {
		"name": "%s",
		"email": "%s"
	  },
	  "langs": [],
	  "tosUrl": "",
	  "repositoryUrl": "https://github.com/laminne/qip",
	  "themeColor": "#8b819a"
	}
  }`,
		name,
		description,
		maintainerName,
		email,
	)
}
