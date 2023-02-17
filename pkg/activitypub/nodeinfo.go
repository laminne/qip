package activitypub

import (
	"fmt"

	"github.com/approvers/qip/pkg/utils/config"
)

func NodeInfo() string {
	return fmt.Sprintf(`{
		"links": [
			{
				"rel": "http://nodeinfo.diaspora.software/ns/schema/2.0",
				"href": "https://%v/nodeinfo/2.0"
			}
		]
	}`,
		config.QipConfig.FQDN)
}

func NodeInfo2() string {
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
	  "outbound": [
		"atom1.0",
		"rss2.0"
	  ]
	},
	"openRegistrations": false,
	"usage": {
	  "users": {
		"total": 1,
		"activeHalfyear": null,
		"activeMonth": null
	  },
	  "localPosts": 0,
	  "localComments": 0
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
	  "repositoryUrl": "https://github.com/approvers/qip",
	  "feedbackUrl": "",
	  "disableRegistration": true,
	  "disableLocalTimeline": false,
	  "disableGlobalTimeline": false,
	  "emailRequiredForSignup": true,
	  "enableHcaptcha": false,
	  "enableRecaptcha": false,
	  "maxNoteTextLength": 3000,
	  "enableEmail": true,
	  "enableServiceWorker": true,
	  "proxyAccountName": "Ghost",
	  "themeColor": "#8b819a"
	}
  }`,
		config.QipConfig.Meta.Name,
		config.QipConfig.Meta.Description,
		config.QipConfig.Meta.Maintainer.Name,
		config.QipConfig.Meta.Maintainer.Email,
	)
}
