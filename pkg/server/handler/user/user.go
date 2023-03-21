package user

import (
	"net/http"

	"github.com/approvers/qip/pkg/controller"
	"github.com/approvers/qip/pkg/errorType"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/server/serverErrors"
	"github.com/approvers/qip/pkg/utils/id"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	controller controller.UserController
}

func NewUserHandler(userRepository repository.UserRepository, fileRepository repository.FileRepository, instanceRepository repository.InstanceRepository) *Handler {
	c := controller.NewUserController(userRepository, fileRepository, instanceRepository)
	return &Handler{controller: *c}
}

func (h *Handler) FindByID(c echo.Context) error {
	uid := c.Param("id")
	res, err := h.controller.FindUserByID(id.SnowFlakeID(uid))
	if err != nil {
		e, code := errorConverter(err)
		return c.JSON(code, e)
	}
	return c.JSON(200, res)
}

func (h *Handler) FindByAcct(c echo.Context) error {
	acct := c.Param("acct")
	res, err := h.controller.FindUserByAcct(acct)
	if err != nil {
		e, code := errorConverter(err)
		return c.JSON(code, e)
	}
	return c.JSON(200, res)
}

func errorConverter(e error) (serverErrors.CommonAPIErrorResponseJSON, int) {
	switch e.(type) {
	case *errorType.ErrNotFound:
		return serverErrors.UserNotFoundErrorResponseJSON, http.StatusNotFound
	case *errorType.ErrExists:
		return serverErrors.InternalErrorResponseJSON, http.StatusConflict
	default:
		return serverErrors.InternalErrorResponseJSON, http.StatusInternalServerError
	}
}
