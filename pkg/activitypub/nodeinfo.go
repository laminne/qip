package activitypub

import "fmt"

const nodeinfo = `{
		"links": [
		  {
			"rel": "http://nodeinfo.diaspora.software/ns/schema/2.0",
			"href": "https://%v/nodeinfo/2.0"
		  }
		]
}`

const nodeinfo2 = `{
	"version": "2.0",
	"software": {
	  "name": "qip",
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
	  "nodeName": "qip Dev",
	  "nodeDescription": "qip",
	  "maintainer": {
		"name": "test",
		"email": "test@example.com"
	  },
	  "langs": [],
	  "tosUrl": "",
	  "repositoryUrl": "",
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
  }`
const ServerFQDN = "np.test.laminne33569.net"

func NodeInfo() string  { return fmt.Sprintf(nodeinfo, ServerFQDN) }
func NodeInfo2() string { return nodeinfo2 }
