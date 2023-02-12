package server

import (
	"fmt"
	"github.com/laminne/notepod/pkg/types"
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

	e.GET("/users/test", userAcctHandler)
	e.GET("/users/1", userAcctHandler)
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

	if acct == "" {
		return c.Blob(http.StatusBadRequest, "plain/text", []byte(""))
	}

	fmt.Println(acct)

	r, err := activitypub.WebFinger(acct)
	if err != nil {
		return c.Blob(http.StatusUnprocessableEntity, "plain/text", []byte(""))
	}

	return c.Blob(http.StatusAccepted, "application/jrd+json; charset=utf-8", []byte(r))
}

func userAcctHandler(c echo.Context) error {
	if strings.Contains(c.Request().Header.Get("Accept"), "application/activity+json") {
		res := activitypub.Person(types.PersonResponseArgs{
			ID:             "1",
			UserName:       "test",
			UserScreenName: "test",
			Summary:        "<p>Hello Fediverse World</p>",
			Icon: struct {
				Url       string
				Sensitive bool
				Name      interface{}
			}{
				Url:       "https://s3.arkjp.net/misskey/a29d961e-9347-469b-b959-5b0c8ae12d8b.png",
				Sensitive: false,
				Name:      nil,
			},
			Image: struct {
				Url       string
				Sensitive bool
				Name      interface{}
			}{
				Url:       "https://s3.arkjp.net/misskey/webpublic-147f04c9-b5c1-4a91-a0cf-28bea9cd267a.png",
				Sensitive: false,
				Name:      nil,
			},
			Tag:                       nil,
			ManuallyApprovesFollowers: false,
			PublicKey:                 "-----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnIizFGc9AR3Mv0x0Gasf\\nTjrPIr7eztYe6xWjWqt4cnIQ/\\npLPR/ZanVZ7v5VGo8jD+X5Y7WXYxhkZrYZg6xWv\\nlcoQxxr07G72btUntWEkXYTSxEeY64C6Qo8Mh+zSdfU9MGAeUyNJS9VhpsS1yvMF\\nlvuTYB9rv1j+CMg0hDui8MEr0ngLkdI+l+mgBLVdVKxyxb7MMLn/24dphINIMPAU\\nFN7piy6EP3nZ6oOCsnFLQqZR+dnYKHueyGuWl++zgglL7aZGaSVXRddcUTmDduTE\\n+uAgd/q6xSiM16DPnIDac7MREsp5wTSaP9jU2618FWV5r2Iljve0ZKnEn+G/Zna2\\nHwIDAQAB\\n-----END PUBLIC KEY-----",
		})

		return c.JSONBlob(200, res)
	}
	return c.String(404, ``)
}
