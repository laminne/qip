# ActivityPub

## nodeinfo

```json
{
  "version": "2.0",
  "software": {
    "name": "実装名",
    "version": "バージョン"
  },
  "protocols": ["activitypub"],
  "services": {
    "inbound": [],
    "outbound": ["atom1.0", "rss2.0"]
  },
  "openRegistrations": false,
  "usage": {
    "users": {
      "total": 0,
      "activeHalfyear": null,
      "activeMonth": null
    },
    "localPosts": 0,
    "localComments": 0
  },
  "metadata": {
    "nodeName": "インスタンス名",
    "nodeDescription": "インスタンス概要",
    "maintainer": {
      "name": "メンテナ名",
      "email": "メアド"
    },
    "langs": [],
    "tosUrl": "利用規約",
    "repositoryUrl": "レポジトリURL",
    "feedbackUrl": "フィードバックURL",
    "disableRegistration": false,
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
}
```
