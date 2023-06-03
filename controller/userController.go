package controller

import (
	"net/http"

	"reglog/models/payload"
	"reglog/usecase"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	id := GetUserId(c)
	user, err := usecase.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get the user",
		"data":    user,
	})
}

func GetUserId(c echo.Context) int {
	user := c.Get("user")
	if user == nil {
		return 0
	}

	userId, ok := user.(int)
	if !ok {
		return 0
	}

	return userId
}

func UpdateUserController(c echo.Context) error {
	var req payload.UpdateUser
	id := GetUserId(c)

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.UpdateProfil(id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

func DeleteUserController(c echo.Context) error {
	id := GetUserId(c)

	err := usecase.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}
