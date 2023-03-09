package router

import (
	"fmt"

	"github.com/approvers/qip/pkg/server/serverErrors"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/server/handler/post"

	"github.com/approvers/qip/pkg/server/handler/user"

	"go.uber.org/zap"

	"github.com/approvers/qip/pkg/repository/dummy"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(port int) {
	userRepository := dummy.NewUserRepository(UserMockData)
	postRepository := dummy.NewPostRepository([]domain.Post{})
	userHandler := user.NewUserHandler(userRepository)
	postHandler := post.NewPostHandler(postRepository)

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
	{
		api.POST("/post", postHandler.Post)

		api.GET("/users/:id", userHandler.FindByID)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func ErrorHandler(err error, c echo.Context) {
	if h, ok := err.(*echo.HTTPError); ok {
		if h.Code == 404 {
			if err := c.JSON(404, serverErrors.NotFoundErrorResponseJSON); err != nil {
				c.Logger().Error(err)
			}
		}
		if h.Code == 503 {
			if err := c.JSON(503, serverErrors.InternalErrorResponseJSON); err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
