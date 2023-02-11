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
	e.GET("/users/acct:test", func(c echo.Context) error {
		fmt.Println("Hello")
		return c.String(404, "not found")
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

func personHandler(c echo.Context) error {

	return c.Blob(http.StatusAccepted, "application/jrd+json; charset=utf-8", []byte(""))
}
