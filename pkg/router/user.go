package router

import (
	"encoding/json"
	"io"

	"github.com/approvers/qip/pkg/controller/models"
	"github.com/labstack/echo/v4"
)

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
