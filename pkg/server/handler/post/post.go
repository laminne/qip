package post

import (
	"fmt"
	"net/http"

	"github.com/approvers/qip/pkg/errorType"

	"github.com/approvers/qip/pkg/server/serverErrors"

	"github.com/approvers/qip/pkg/controller"
	"github.com/approvers/qip/pkg/controller/models"
	"github.com/approvers/qip/pkg/repository"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	controller controller.PostController
}

func NewPostHandler(repo repository.PostRepository) *Handler {
	return &Handler{controller: *controller.NewPostController(repo)}
}

func (h *Handler) Post(c echo.Context) error {
	req := models.CreatePostRequestJSON{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, serverErrors.InvalidRequestErrorResponseJSON)
	}

	// ToDo: Authorを正しく指定する
	res, err := h.controller.Create(req.Body, "123", req.Visibility)
	if err != nil {
		e, code := errorConverter(err)
		return c.JSON(code, e)
	}
	fmt.Println(res)
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) FindByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.controller.FindByID(id)
	if err != nil {
		e, code := errorConverter(err)
		return c.JSON(code, e)
	}

	return c.JSON(http.StatusOK, res)
}

func errorConverter(e error) (serverErrors.CommonAPIErrorResponseJSON, int) {
	switch e.(type) {
	case *errorType.ErrNotFound:
		return serverErrors.PostNotFoundErrorResponseJSON, http.StatusNotFound
	case *errorType.ErrExists:
		return serverErrors.InternalErrorResponseJSON, http.StatusConflict
	default:
		return serverErrors.InternalErrorResponseJSON, http.StatusInternalServerError
	}
}
