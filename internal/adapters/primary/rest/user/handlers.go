package user

import (
	"cleanArchitecture/internal/application/dto"
	"cleanArchitecture/internal/application/interactors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handlers struct {
	interactor *interactors.UserInteractor
}

func RegisterUserHandlers(e *echo.Echo, interactor *interactors.UserInteractor) *Handlers {
	handlers := &Handlers{
		interactor: interactor,
	}

	routeGroup := e.Group("/users")
	{
		routeGroup.POST("", handlers.CreateUser)
		routeGroup.GET("", handlers.GetUsers)
	}

	return handlers
}

func (h *Handlers) CreateUser(c echo.Context) error {

	var req CreateUserRq
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.interactor.CreateUser(dto.CreateUserRq{
		Name: req.Name,
		Age:  req.Age,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusCreated)
}

func (h *Handlers) GetUsers(c echo.Context) error {
	users, _ := h.interactor.GetAll()
	resp := make([]UserResponse, 0)
	for _, user := range users {
		resp = append(resp, UserResponse{
			Name: user.Name,
			Age:  user.Age,
			Id:   user.Id.String(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}
