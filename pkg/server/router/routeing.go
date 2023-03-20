package router

import "github.com/labstack/echo/v4"

func rootRouter(e *echo.Echo) {
	e.GET("/.well-known/nodeinfo", apHandler.GetNodeInfo)
	e.GET("/nodeinfo/2.0", apHandler.GetNodeInfo2)

	e.POST("/api/v1/login", authHandler.LoginHandler)
	api := e.Group("/api/v1")
	{
		api.Use(authHandler.TokenMiddlewareHandlerFunc)
		api.POST("/posts", postHandler.Post)
		api.GET("/posts/:id", postHandler.FindByID)

		api.GET("/users/:id", userHandler.FindByID)
	}
}
