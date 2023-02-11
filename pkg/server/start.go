package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laminne/notepod/pkg/activitypub"
)

func StartServer(port int) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloHandler)
	e.GET("/.well-known/nodeinfo", nodeInfoHandler)
	e.GET("/nodeinfo/2.0", nodeInfo2Handler)
	e.GET("/.well-known/webfinger", webFingerHandler)

	e.GET("/users/test", func(c echo.Context) error {
		if strings.Contains(c.Request().Header.Get("Accept"), "application/activity+json") {
			return c.String(200, `{
				"@context": [
				  "https://www.w3.org/ns/activitystreams",
				  "https://w3id.org/security/v1",
				  {
					"manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
					"sensitive": "as:sensitive",
					"Hashtag": "as:Hashtag",
					"quoteUrl": "as:quoteUrl",
					"toot": "http://joinmastodon.org/ns#",
					"Emoji": "toot:Emoji",
					"featured": "toot:featured",
					"discoverable": "toot:discoverable",
					"schema": "http://schema.org#",
					"PropertyValue": "schema:PropertyValue",
					"value": "schema:value",
					"misskey": "https://misskey-hub.net/ns#",
					"_misskey_content": "misskey:_misskey_content",
					"_misskey_quote": "misskey:_misskey_quote",
					"_misskey_reaction": "misskey:_misskey_reaction",
					"_misskey_votes": "misskey:_misskey_votes",
					"_misskey_talk": "misskey:_misskey_talk",
					"isCat": "misskey:isCat",
					"vcard": "http://www.w3.org/2006/vcard/ns#"
				  }
				],
				"type": "Person",
				"id": "https://np.test.laminne33569.net/users/1",
				"inbox": "https://np.test.laminne33569.net/inbox",
				"outbox": "https://np.test.laminne33569.net/users/1/outbox",
				"followers": "https://np.test.laminne33569.net/users/1/followers",
				"following": "https://np.test.laminne33569.net/users/1/following",
				"featured": "https://np.test.laminne33569.net/users/1/collections/featured",
				"sharedInbox": "https://np.test.laminne33569.net/inbox",
				"endpoints": {
				  "sharedInbox": "https://np.test.laminne33569.net/inbox"
				},
				"url": "https://np.test.laminne33569.net/@test",
				"preferredUsername": "test",
				"name": "test",
				"summary": "<p>Hello Fediverse World</p>",
				"icon": {
				  "type": "Image",
				  "url": "https://s3.arkjp.net/misskey/a29d961e-9347-469b-b959-5b0c8ae12d8b.png",
				  "sensitive": false,
				  "name": null
				},
				"image": {
				  "type": "Image",
				  "url": "https://s3.arkjp.net/misskey/webpublic-147f04c9-b5c1-4a91-a0cf-28bea9cd267a.png",
				  "sensitive": false,
				  "name": null
				},
				"tag": [],
				"manuallyApprovesFollowers": false,
				"discoverable": true,
				"publicKey": {
				  "id": "https://np.test.laminne33569.net/users/1#main-key",
				  "type": "Key",
				  "owner": "https://np.test.laminne33569.net/users/1",
				  "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnIizFGc9AR3Mv0x0Gasf\nTjrPIr7eztYe6xWjWqt4cnIQ/\npLPR/ZanVZ7v5VGo8jD+X5Y7WXYxhkZrYZg6xWv\nlcoQxxr07G72btUntWEkXYTSxEeY64C6Qo8Mh+zSdfU9MGAeUyNJS9VhpsS1yvMF\nlvuTYB9rv1j+CMg0hDui8MEr0ngLkdI+l+mgBLVdVKxyxb7MMLn/24dphINIMPAU\nFN7piy6EP3nZ6oOCsnFLQqZR+dnYKHueyGuWl++zgglL7aZGaSVXRddcUTmDduTE\n+uAgd/q6xSiM16DPnIDac7MREsp5wTSaP9jU2618FWV5r2Iljve0ZKnEn+G/Zna2\nHwIDAQAB\n-----END PUBLIC KEY-----"
				},
				"isCat": false,
				"vcard:bday": "2023-02-21",
				"vcard:Address": "Test"
			  }
			  `)
		}
		return c.String(404, ``)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func nodeInfoHandler(c echo.Context) error {
	nodeinfo := activitypub.NodeInfo()

	return c.Blob(http.StatusAccepted, "application/json+activity", []byte(nodeinfo))
}

func nodeInfo2Handler(c echo.Context) error {
	data := activitypub.NodeInfo2()
	return c.Blob(http.StatusAccepted, "application/json+activity", []byte(data))
}

func webFingerHandler(c echo.Context) error {
	acct := c.QueryParam("resource")
	if len(strings.Split(acct, ":")) != 0 {
		acct = acct[5:]
	}
	fmt.Println(acct)
	if acct == "" {
		return c.Blob(http.StatusBadRequest, "plain/text", []byte(""))
	}

	r := activitypub.WebFinger(acct)

	return c.Blob(http.StatusAccepted, "application/jrd+json; charset=utf-8", []byte(r))
}
