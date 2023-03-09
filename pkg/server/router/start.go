package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/approvers/qip/pkg/utils/logger"

	"github.com/approvers/qip/pkg/server/serverErrors"

	"github.com/approvers/qip/pkg/server/handler/post"

	"github.com/approvers/qip/pkg/server/handler/user"

	"github.com/approvers/qip/pkg/repository/dummy"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(port int) {
	zapLogger := logger.NewZapLogger(nil)
	userRepository := dummy.NewUserRepository(UserMockData)
	postRepository := dummy.NewPostRepository(PostMockData)
	userHandler := user.NewUserHandler(userRepository)
	postHandler := post.NewPostHandler(postRepository)

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogUserAgent: true,
		LogMethod:    true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			zapLogger.Debug(fmt.Sprintf("[API] %v %v %v %v %v", v.URI, v.Status, v.Latency, v.Method, v.UserAgent))
			return nil
		},
	}))

	e.Use(middleware.Recover())
	e.HTTPErrorHandler = ErrorHandler

	api := e.Group("/api/v1")
	{
		api.POST("/posts", postHandler.Post)
		api.GET("/posts/:id", postHandler.FindByID)

		api.GET("/users/:id", userHandler.FindByID)
	}

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
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
