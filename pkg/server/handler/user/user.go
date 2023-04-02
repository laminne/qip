package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/laminne/qip/pkg/controller"
	"github.com/laminne/qip/pkg/errorType"
	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/server/serverErrors"
	"github.com/laminne/qip/pkg/utils/id"
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
