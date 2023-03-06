package router

import (
	"github.com/approvers/qip/pkg/controller"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	controller controller.UserController
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	c := controller.NewUserController(userRepository)
	return &UserHandler{controller: *c}
}

func (h *UserHandler) findUserByIDHandler(c echo.Context) error {
	uid := c.Param("id")
	res, err := h.controller.FindUserByID(id.SnowFlakeID(uid))
	if err != nil {
		return c.String(500, "Internal error")
	}
	return c.JSON(200, res)
}
