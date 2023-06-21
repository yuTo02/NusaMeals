package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	UserUseCase usecase.UserUseCase
}

func NewAuthController(uc usecase.UserUseCase) *AuthController {
	return &AuthController{
		UserUseCase: uc,
	}
}

func (h *AuthController) RegisterUserController(c echo.Context) error {
	var requestRegister request.RegisterUser
	if err := c.Bind(&requestRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := c.Validate(requestRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	//_, e := c.Cookie("JWTCookie")
	//
	//if e == nil {
	//	return echo.NewHTTPError(http.StatusForbidden, "Already logged in")
	//}

	if requestRegister.Password != requestRegister.RetypePassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}

	requestRegister.Role = "user"
	err := h.UserUseCase.RegisterUser(requestRegister)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully create your account",
	})
}

func (h *AuthController) RegisterAdminController(c echo.Context) error {
	var requestRegister request.RegisterUser
	if err := c.Bind(&requestRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := c.Validate(requestRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	//_, e := c.Cookie("JWTCookie")
	//
	//if e == nil {
	//	return echo.NewHTTPError(http.StatusForbidden, "Already logged in")
	//}

	if requestRegister.Password != requestRegister.RetypePassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}

	requestRegister.Role = "admin"
	err := h.UserUseCase.RegisterUser(requestRegister)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully create your account",
	})
}

func (h *AuthController) LoginController(c echo.Context) error {
	var requestLogin request.LoginUser
	if err := c.Bind(&requestLogin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := c.Validate(requestLogin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	//if _, err := c.Cookie("JWTCookie"); err == nil {
	//	return echo.NewHTTPError(http.StatusForbidden, "Already logged in")
	//}

	responseLogin, err := h.UserUseCase.LoginUser(requestLogin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//token, err := middleware.CreateToken(user.Username, user.Role, user.ID)
	//if err != nil {
	//	return err
	//}
	//cookie.CreateJWTCookies(c, token)

	return c.JSON(http.StatusOK, responseLogin)
}
func (h *AuthController) LogoutController(c echo.Context) error {
	// Get the JWT token from the request headers or cookies
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		cookie, err := c.Cookie("JWTCookie")
		if err == nil {
			tokenString = cookie.Value
		}
	}

	// Parse and validate the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Specify the key or secret used to sign the JWT token
		return []byte("SECRET123"), nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		// Extract user ID from the claims
		userID, err := strconv.ParseUint(claims.Subject, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse user ID")
		}

		// Perform logout logic in the user use case
		err = h.UserUseCase.LogoutUser(uint(userID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to logout user")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully logged out",
		})
	}

	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
}

//func LogoutController(c echo.Context) error {
//	cookie, err := c.Cookie("JWTCookie")
//	if err != nil {
//		return echo.NewHTTPError(http.StatusNotAcceptable, "Not logged in yet")
//	}
//	cookie.Expires = time.Now().Add(-1 * time.Hour)
//	c.SetCookie(cookie)
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "success to logout",
//	})
//}

//func Authorization(c echo.Context) (string, uint) {
//	cookie, _ := c.Cookie("JWTCookie")
//	token, _ := jwt.Parse(cookie.Value, nil)
//	claims, _ := token.Claims.(jwt.MapClaims)
//	username := claims["username"].(string)
//	id := uint(claims["user_id"].(float64))
//	return username, id
//}
