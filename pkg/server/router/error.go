package router

import (
	"github.com/approvers/qip/pkg/server/serverErrors"
	"github.com/labstack/echo/v4"
)

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
