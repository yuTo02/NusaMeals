package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
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

func (h *UserController) GetUserByUsername(c echo.Context) error {
	var Username string = c.Param("username")

	user, err := h.UserUseCase.GetUserByUsername(Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserController) GetUserByEmail(c echo.Context) error {
	var Email string = c.Param("email")

	user, err := h.UserUseCase.GetUserByEmail(Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserController) UpdateUser(c echo.Context) error {
	ID := c.Param("id")

	// Parse request body to get update data
	var updateUser request.UpdateUser
	if err := c.Bind(&updateUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Convert ID to uint
	idUint, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	userID := uint(idUint)

	// Update user
	updatedUser, err := h.UserUseCase.UpdateUser(userID, updateUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}
