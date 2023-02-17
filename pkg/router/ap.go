package router

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/approvers/qip/pkg/activitypub"
	"github.com/labstack/echo/v4"
)

func nodeInfoHandler(c echo.Context) error {
	nodeinfo := activitypub.NodeInfo()

	return c.Blob(http.StatusAccepted, ActivityJSONLDContentsType, []byte(nodeinfo))
}

func nodeInfo2Handler(c echo.Context) error {
	data := activitypub.NodeInfo2()
	return c.Blob(http.StatusAccepted, ActivityJSONLDContentsType, []byte(data))
}

func webFingerHandler(c echo.Context) error {
	acct := c.QueryParam("resource")

	if acct == "" {
		return c.Blob(http.StatusBadRequest, PlainTextContentsType, []byte(""))
	}

	r, err := activitypub.WebFinger(acct)
	if err != nil {
		return c.Blob(http.StatusUnprocessableEntity, PlainTextContentsType, []byte(""))
	}

	return c.Blob(http.StatusAccepted, JSONLDContentsType, []byte(r))
}

func userAcctHandler(c echo.Context) error {
	if strings.Contains(c.Request().Header.Get("Accept"), "application/activity+json") {
		param := c.Param("name")
		name := param

		if len(param) == 0 {
			return c.String(404, "")
		}
		if string(param[0]) == "@" {
			name = param[1:]
		} else if string(param[:5]) == "acct:" {
			name = param[5:]
		}

		res := apController.GetUser(name)
		j, _ := json.Marshal(res)
		return c.JSONBlob(200, j)
	}
	return c.String(404, "")
}
