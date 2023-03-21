package follow

import (
	"net/http"

	"github.com/approvers/qip/pkg/utils/token"

	"github.com/approvers/qip/pkg/controller"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/server/serverErrors"
	"github.com/approvers/qip/pkg/utils/id"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	controller  controller.FollowController
	tokenParser token.JWTTokenParser
}

func NewFollowHandler(repo repository.FollowRepository, key string) *Handler {
	return &Handler{controller: *controller.NewFollowController(repo), tokenParser: *token.NewJWTTokenParser(key)}
}

func (h *Handler) Create(c echo.Context) error {
	target := c.Param("id")
	uid, err := h.tokenParser.Parse(target)
	if err != nil {
		return c.JSON(http.StatusBadRequest, serverErrors.InvalidRequestErrorResponseJSON)
	}
	err = h.controller.Create(uid, id.SnowFlakeID(target))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, serverErrors.InternalErrorResponseJSON)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) FindUserFollow(c echo.Context) error {
	i := c.Param("id")
	follower, err := h.controller.FindUserFollow(id.SnowFlakeID(i))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, serverErrors.InternalErrorResponseJSON)
	}
	return c.JSON(http.StatusOK, follower)
}

func (h *Handler) FindUserFollower(c echo.Context) error {
	i := c.Param("id")
	follower, err := h.controller.FindUserFollower(id.SnowFlakeID(i))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, serverErrors.InternalErrorResponseJSON)
	}
	return c.JSON(http.StatusOK, follower)
}
