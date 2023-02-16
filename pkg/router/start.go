package router

import (
	"database/sql"
	"fmt"

	"github.com/approvers/qip/pkg/controller"

	"github.com/approvers/qip/pkg/repository"

	bun2 "github.com/approvers/qip/pkg/repository/bun"

	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/uptrace/bun"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var UserRepository repository.UserRepository
var apController controller.ActivityPubController
var userController controller.UserController

func StartServer(port int) {
	// ToDo: ここのハードコーディングをやめる
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

	wk := e.Group("/.well-known")
	wk.GET("/nodeinfo", nodeInfoHandler)
	wk.GET("/webfinger", webFingerHandler)

	e.GET("/nodeinfo/2.0", nodeInfo2Handler)
	e.GET("/users/:name", userAcctHandler)

	api := e.Group("/api")
	api.POST("/users", createUserHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
