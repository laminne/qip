package user

import (
	"github.com/approvers/qip/pkg/controller"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	controller controller.UserController
}

func NewUserHandler(userRepository repository.UserRepository) *Handler {
	c := controller.NewUserController(userRepository)
	return &Handler{controller: *c}
}

func (h *Handler) FindByID(c echo.Context) error {
	uid := c.Param("id")
	res, err := h.controller.FindUserByID(id.SnowFlakeID(uid))
	if err != nil {
		return c.String(500, "Internal error")
	}
	return c.JSON(200, res)
}
