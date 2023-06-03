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
	productRepository := repository.NewProductRepository(cfg.DBConn)

	userUseCase := usecase.NewUserUseCase(userRepository, cfg.JwtProvider)
	productUseCase := usecase.NewProductUseCase(productRepository)

	// Routes

	// AUTH
	authController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/register", authController.RegisterUserController)
	cfg.Echo.POST("/login", authController.LoginController)

	adminController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/admin/register", adminController.RegisterAdminController)

	// USER
	userController := controller.NewUserController(userUseCase)
	user := cfg.Echo.Group("/users", authMiddleware.IsAuthenticated())
	user.GET("", userController.GetAllUser, authMiddleware.IsAdmin)
	user.GET("/:id", userController.GetUserByID, authMiddleware.IsUser)

	// PRODUCT
	productController := controller.NewProductController(productUseCase)
	product := cfg.Echo.Group("/products", authMiddleware.IsAuthenticated())
	product.POST("", productController.CreateProductController, authMiddleware.IsAdmin)

}
