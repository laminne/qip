package post

import (
	"net/http"

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
		return c.JSON(http.StatusInternalServerError, serverErrors.InternalErrorResponseJSON)
	}

	return c.JSON(http.StatusOK, res)
}
