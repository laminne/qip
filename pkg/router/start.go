package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/approvers/qip/pkg/controller/models"

	"github.com/approvers/qip/pkg/controller"

	"github.com/approvers/qip/pkg/repository"

	bun2 "github.com/approvers/qip/pkg/repository/bun"

	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/uptrace/bun"

	"github.com/approvers/qip/pkg/activitypub"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var UserRepository repository.UserRepository
var apController controller.ActivityPubController
var userController controller.UserController

func StartServer(port int) {
	db := bun.NewDB(
		sql.OpenDB(
			pgdriver.NewConnector(
				pgdriver.WithDSN("postgres://postgres:qip@localhost:5432/qip?sslmode=disable"),
			),
		),
		pgdialect.New(),
	)

	UserRepository = *bun2.NewUserRepository(db)
	apController = *controller.NewActivityPubController(UserRepository)
	userController = *controller.NewUserController(UserRepository)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.GET("/", helloHandler)
	e.GET("/.well-known/nodeinfo", nodeInfoHandler)
	e.GET("/nodeinfo/2.0", nodeInfo2Handler)
	e.GET("/.well-known/webfinger", webFingerHandler)
	e.GET("/users/:name", userAcctHandler)

	e.POST("/api/users", createUserHandler)

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
			name = param[1:]
		} else if string(param[:5]) == "acct:" {
			name = param[5:]
		}

		res := apController.GetUser(name)
		j, _ := json.Marshal(res)
		return c.JSONBlob(200, j)
	}
	return c.String(404, ``)
}

func createUserHandler(c echo.Context) error {
	b := models.CreateUserRequestJSON{}
	body, _ := io.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &b)
	if err != nil {
		return err
	}

	res, err := userController.CreateUser(b)
	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
