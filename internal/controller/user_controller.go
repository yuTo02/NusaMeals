package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reglog/internal/usecase"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: uc,
	}
}

func (h *UserController) GetAllUser(c echo.Context) error {
	users, err := h.UserUseCase.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserController) GetUserByID(c echo.Context) error {
	var ID string = c.Param("id")

	user, err := h.UserUseCase.GetUserByID(ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
