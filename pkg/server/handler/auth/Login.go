package auth

import (
	"net/http"
	"strings"

	"github.com/approvers/qip/pkg/repository"

	"github.com/approvers/qip/pkg/controller"
	"github.com/approvers/qip/pkg/controller/models"
	"github.com/approvers/qip/pkg/server/serverErrors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	authController controller.AuthController
}

func NewHandler(r repository.UserRepository, key string) *Handler {
	return &Handler{authController: *controller.NewAuthController(r, key)}
}

func (t *Handler) TokenMiddlewareHandlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// トークンを取り出す
		h := c.Request().Header.Get("authorization")
		if len(h) < 7 {
			return c.JSON(http.StatusUnauthorized, serverErrors.UnAuthorizedRequestErrorResponseJSON)
		}
		token := strings.Split(h, " ")[1]

		// 検証する
		if t.authController.CheckToken(token) {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, serverErrors.UnAuthorizedRequestErrorResponseJSON)
	}
}

func (t *Handler) LoginHandler(c echo.Context) error {
	req := models.LoginRequestJSON{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, serverErrors.InvalidRequestErrorResponseJSON)
	}

	token, err := t.authController.Login(req.Name, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, serverErrors.InternalErrorResponseJSON)
	}

	return c.JSON(http.StatusOK, models.LoginResponseJSON{Token: token})
}
