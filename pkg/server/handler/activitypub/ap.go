package activitypub

import (
	"net/http"

	"github.com/approvers/qip/pkg/controller"
	"github.com/labstack/echo/v4"
)

type ApHandler struct {
	controller controller.ActivityPubController
}

func NewApHandler() *ApHandler {
	return &ApHandler{}
}

func (h *ApHandler) GetNodeInfo(c echo.Context) error {
	return c.Blob(http.StatusOK, "application/json", []byte(h.controller.GetNodeInfo()))
}

func (h *ApHandler) GetNodeInfo2(c echo.Context) error {
	return c.Blob(http.StatusOK, "application/json", []byte(h.controller.GetNodeInfo2()))
}
