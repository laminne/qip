package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/laminne/notepod/pkg/controller"

	"github.com/laminne/notepod/pkg/repository"

	bun2 "github.com/laminne/notepod/pkg/repository/bun"

	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/uptrace/bun"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laminne/notepod/pkg/activitypub"
)

var UserRepository repository.UserRepository
var apController controller.ActivityPubController

func StartServer(port int) {
	db := bun.NewDB(
		sql.OpenDB(
			pgdriver.NewConnector(
				pgdriver.WithDSN("postgres://postgres:notepod@localhost:5432/notepod?sslmode=disable"),
			),
		),
		pgdialect.New(),
	)
	UserRepository = *bun2.NewUserRepository(db)
	apController = *controller.NewActivityPubController(UserRepository)

	e := echo.New()

	a, _ := os.OpenFile("notepod.log", os.O_RDWR|os.O_CREATE, 0660)
	defer a.Close()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: a,
	}))
	//e.Use(middleware.Recover())

	e.GET("/", helloHandler)
	e.GET("/.well-known/nodeinfo", nodeInfoHandler)
	e.GET("/nodeinfo/2.0", nodeInfo2Handler)
	e.GET("/.well-known/webfinger", webFingerHandler)

	e.GET("/users/:name", userAcctHandler)

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

	r, err := activitypub.WebFinger(acct)
	if err != nil {
		return c.Blob(http.StatusUnprocessableEntity, "plain/text", []byte(""))
	}

	return c.Blob(http.StatusAccepted, "application/jrd+json; charset=utf-8", []byte(r))
}

func userAcctHandler(c echo.Context) error {
	if strings.Contains(c.Request().Header.Get("Accept"), "application/activity+json") {
		param := c.Param("name")
		name := param

		if len(param) == 0 {
			return c.String(404, "")
		}
		if string(param[0]) == "@" {
			name = string(param[1:])
		} else if string(param[:5]) == "acct:" {
			name = string(param[5:])

		}

		res := apController.GetUser(name)
		j, _ := json.Marshal(res)
		return c.JSONBlob(200, j)
	}
	return c.String(404, ``)
}
