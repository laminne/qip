package post

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/approvers/qip/pkg/utils/token"

	"github.com/approvers/qip/pkg/errorType"

	"github.com/approvers/qip/pkg/server/serverErrors"

	"github.com/approvers/qip/pkg/controller"
	"github.com/approvers/qip/pkg/controller/models"
	"github.com/approvers/qip/pkg/repository"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	controller  controller.PostController
	tokenParser token.JWTTokenParser
}

func NewPostHandler(repo repository.PostRepository, key string, userRepository repository.UserRepository) *Handler {
	return &Handler{
		controller:  *controller.NewPostController(repo, userRepository),
		tokenParser: *token.NewJWTTokenParser(key),
	}
}

func (h *Handler) Post(c echo.Context) error {
	req := models.CreatePostRequestJSON{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, serverErrors.InvalidRequestErrorResponseJSON)
	}
	token := strings.Split(c.Request().Header.Get("authorization"), " ")[1]
	uID, err := h.tokenParser.Parse(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, serverErrors.InternalErrorResponseJSON)
	}

	res, err := h.controller.Create(req.Body, uID, req.Visibility)
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

func (h *Handler) FindByAuthor(c echo.Context) error {
	id := c.Param("id")
	res, err := h.controller.FindByAuthorID(id)
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
