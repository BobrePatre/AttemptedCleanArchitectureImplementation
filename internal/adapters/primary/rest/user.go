package rest

import (
	"cleanArchitecture/internal/application/interactors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandlers struct {
	interactor *interactors.UserInteractor
}

func RegisterUserHandlers(e *echo.Echo, interactor *interactors.UserInteractor) *UserHandlers {
	handlers := &UserHandlers{
		interactor: interactor,
	}

	routeGroup := e.Group("/users")
	{
		routeGroup.POST("", handlers.CreateUser)
	}

	return handlers
}

func (h *UserHandlers) CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Counter: %d", h.interactor.Incr()))
}
