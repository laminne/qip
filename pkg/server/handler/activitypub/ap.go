package activitypub

import (
	"net/http"

	"github.com/laminne/qip/pkg/repository"

	"github.com/labstack/echo/v4"
	"github.com/laminne/qip/pkg/controller"
)

type ApHandler struct {
	controller controller.ActivityPubController
}

func NewApHandler(r repository.UserRepository, f repository.FileRepository) *ApHandler {
	return &ApHandler{*controller.NewActivityPubController(r, f)}
}

func (h *ApHandler) GetNodeInfo(c echo.Context) error {
	return c.Blob(http.StatusOK, "application/json", []byte(h.controller.GetNodeInfo()))
}

func (h *ApHandler) GetNodeInfo2(c echo.Context) error {
	return c.Blob(http.StatusOK, "application/json", []byte(h.controller.GetNodeInfo2()))
}

func (h *ApHandler) GetWebFinger(c echo.Context) error {
	q := c.QueryParam("resource")
	webFinger, err := h.controller.GetWebFinger(q)
	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	return c.Blob(http.StatusOK, "application/json", []byte(webFinger))
}

func (h *ApHandler) GetPerson(c echo.Context) error {
	i := c.Param("id")
	// ToDo: Acceptするtypeを調べる(application/activity+jsonは動作しない
	person, err := h.controller.GetPerson(i)
	if err != nil {
		return c.Blob(http.StatusInternalServerError, "text/plain", []byte("server error"))
	}
	return c.Blob(http.StatusOK, "application/activity+json", []byte(person))
}
