package router

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/approvers/qip/pkg/repository/dummy"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(port int) {
	repo := dummy.NewUserRepository(UserMockData)
	userHandler := NewUserHandler(repo)

	e := echo.New()

	logger, _ := zap.NewDevelopment()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogUserAgent: true,
		LogMethod:    true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
				zap.Durationp("Latency", &v.Latency),
				zap.String("method", v.Method),
				zap.String("ua", v.UserAgent),
			)

			return nil
		},
	}))

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
