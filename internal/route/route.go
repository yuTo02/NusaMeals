package route

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"reglog/internal/common/middleware"
	"reglog/internal/controller"
	"reglog/internal/repository"
	"reglog/internal/usecase"
)

type Config struct {
	// ECHO TOP LEVEL INSTANCE
	Echo        *echo.Echo
	DBConn      *gorm.DB
	JwtProvider *middleware.JWTProvider
}

func (cfg *Config) New() {
	// Get Auth middleware to filter authorization user/admin
	authMiddleware := middleware.NewAuthMiddleware(cfg.JwtProvider)

	// dependency injection
	userRepository := repository.NewUserRepository(cfg.DBConn)
	userUseCase := usecase.NewUserUseCase(userRepository, cfg.JwtProvider)
	authController := controller.NewAuthController(userUseCase)
	userController := controller.NewUserController(userUseCase)

	// Routes

	// AUTH
	cfg.Echo.POST("/register", authController.RegisterController)
	cfg.Echo.POST("/login", authController.LoginController)
	// USER
	user := cfg.Echo.Group("users", authMiddleware.IsAuthenticated())
	user.GET("", userController.GetAllUser, authMiddleware.IsAdmin)
	user.GET("/:id", userController.GetAllUser, authMiddleware.IsUser)

}
