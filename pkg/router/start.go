package router

import (
	"fmt"

	"github.com/approvers/qip/pkg/repository/dummy"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(port int) {
	repo := dummy.NewUserRepository(UserMockData)
	userHandler := NewUserHandler(repo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = ErrorHandler

	api := e.Group("/api/v1")
	api.GET("/users/:id", userHandler.findUserByIDHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func ErrorHandler(err error, c echo.Context) {
	if h, ok := err.(*echo.HTTPError); ok {
		if h.Code == 404 {
			c.JSON(404, notFoundErrorResponseJSON)
		}
		if h.Code == 503 {
			c.JSON(503, internalErrorResponseJSON)
		}
	}
}
