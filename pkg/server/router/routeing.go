package router

import "github.com/labstack/echo/v4"

func rootRouter(e *echo.Echo) {
	wk := e.Group("/.well-known")
	{
		wk.GET("/nodeinfo", apHandler.GetNodeInfo)
		wk.GET("/webfinger", apHandler.GetWebFinger)
	}
	e.GET("/nodeinfo/2.0", apHandler.GetNodeInfo2)
	e.GET("/users/:id", apHandler.GetPerson)

	e.POST("/api/v1/login", authHandler.LoginHandler)
	api := e.Group("/api/v1")
	{
		api.Use(authHandler.TokenMiddlewareHandlerFunc)
		api.POST("/posts", postHandler.Post)
		api.GET("/posts/:id", postHandler.FindByID)

		api.POST("/users/:id/follow", followHandler.Create)
		api.GET("/users/:id/follow", followHandler.FindUserFollow)
		api.GET("/users/:id/follower", followHandler.FindUserFollower)
		api.GET("/users/:id", userHandler.FindByID)
	}
}
