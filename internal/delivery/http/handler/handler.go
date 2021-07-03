package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sew810i9/opus/internal/domain/issue"
	"github.com/sew810i9/opus/internal/domain/user"
	"net/http"
)

type Handler struct {
	userService  user.Service
	issueService issue.Service
}

func NewHandler(userService user.Service, issueService issue.Service) *Handler {
	return &Handler{
		userService:  userService,
		issueService: issueService,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	return e
}
