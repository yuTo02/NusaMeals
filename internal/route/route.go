package route

import (
	"reglog/internal/common/middleware"
	"reglog/internal/controller"
	"reglog/internal/repository"
	"reglog/internal/usecase"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
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
	//register ADMIN
	adminController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/admin/register", adminController.RegisterAdminController)

	// USER
	userController := controller.NewUserController(userUseCase)
	user := cfg.Echo.Group("/users", authMiddleware.IsAuthenticated())
	user.GET("/:id", userController.GetUserByID, authMiddleware.IsUser)
	user.GET("/:username", userController.GetUserByUsername, authMiddleware.IsUser)
	user.GET("/:email", userController.GetUserByEmail, authMiddleware.IsUser)
	user.PUT("/:id", userController.UpdateUser, authMiddleware.IsUser)

	//ADMIN
	user.GET("", userController.GetAllUser, authMiddleware.IsAdmin)
	user.GET("/:id", userController.GetUserByID, authMiddleware.IsAdmin)
	user.GET("/:username", userController.GetUserByUsername, authMiddleware.IsAdmin)
	user.GET("/:email", userController.GetUserByEmail, authMiddleware.IsAdmin)
	user.PUT("/:id", userController.UpdateUser, authMiddleware.IsAdmin)

	// PRODUCT
	productController := controller.NewProductController(productUseCase)
	product := cfg.Echo.Group("/products", authMiddleware.IsAuthenticated())
	product.POST("", productController.CreateProductController, authMiddleware.IsAdmin)

}
