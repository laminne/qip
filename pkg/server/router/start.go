package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/approvers/qip/pkg/server/handler/follow"

	"github.com/approvers/qip/pkg/server/handler/activitypub"

	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/approvers/qip/pkg/repository/gormRepository"
	"github.com/approvers/qip/pkg/utils/config"
	"github.com/approvers/qip/pkg/utils/token"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/server/handler/auth"

	"go.uber.org/zap"

	"github.com/approvers/qip/pkg/server/handler/post"

	"github.com/approvers/qip/pkg/server/handler/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	userRepository     repository.UserRepository
	postRepository     repository.PostRepository
	fileRepository     repository.FileRepository
	instanceRepository repository.InstanceRepository
	followRepository   repository.FollowRepository
	userHandler        *user.Handler
	postHandler        *post.Handler
	authHandler        *auth.Handler
	apHandler          *activitypub.ApHandler
	followHandler      *follow.Handler
)

func initServer() {
	if config.QipConfig.Mode != "development" {
		db, err := gorm.Open(postgres.Open(
			fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
				config.QipConfig.DB.User,
				config.QipConfig.DB.Password,
				config.QipConfig.DB.Host,
				config.QipConfig.DB.Port,
				config.QipConfig.DB.DBName,
			),
		),
			&gorm.Config{},
		)
		if err != nil {
			panic("failed to connect database")
		}

		fmt.Printf("[Root] Successfully connect to database\n\tDatabase User: %s\n\tDatabase Host: %s\n\tDatabase Post: %s\n\tDatabase Name: %s\n",
			config.QipConfig.DB.User,
			config.QipConfig.DB.Host,
			config.QipConfig.DB.Port,
			config.QipConfig.DB.DBName)
		userRepository = gormRepository.NewUserRepository(db)
		postRepository = gormRepository.NewPostRepository(db)
		fileRepository = gormRepository.NewFileRepository(db)
		instanceRepository = gormRepository.NewInstanceRepository(db)
		followRepository = gormRepository.NewFollowRepository(db)
	} else {
		userRepository = dummy.NewUserRepository(UserMockData)
		postRepository = dummy.NewPostRepository(PostMockData)
	}
	key := token.SecureRandom(512)
	userHandler = user.NewUserHandler(userRepository, fileRepository, instanceRepository)
	postHandler = post.NewPostHandler(postRepository, key, userRepository)
	authHandler = auth.NewHandler(userRepository, key)
	apHandler = activitypub.NewApHandler(userRepository, fileRepository)
	followHandler = follow.NewFollowHandler(followRepository, key)
}

func StartServer(port int) {
	initServer()
	e := echo.New()
	e.Use(middleware.CORS())

	logger, _ := zap.NewDevelopment()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogUserAgent: true,
		LogMethod:    true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Sugar().Infof("[API] %v %v %v %v %v", v.URI, v.Status, v.Latency, v.Method, v.UserAgent)
			return nil
		},
	}))

	e.Use(middleware.Recover())
	e.HTTPErrorHandler = ErrorHandler
	rootRouter(e)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down server", err)
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
