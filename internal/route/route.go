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
	menuRepository := repository.NewMenuRepository(cfg.DBConn)

	userUseCase := usecase.NewUserUseCase(userRepository, cfg.JwtProvider)
	productUseCase := usecase.NewProductUseCase(productRepository)
	menuUseCase := usecase.NewMenuUseCase(menuRepository)

	// Routes

	// AUTH
	authController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/register", authController.RegisterUserController)
	cfg.Echo.POST("/login", authController.LoginController)
	cfg.Echo.POST("/logout", authController.LogoutController)
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
	admin := cfg.Echo.Group("/admin", authMiddleware.IsAdmin)
	admin.GET("/users", userController.GetAllUser)
	admin.GET("/users/:id", userController.GetUserByID)
	admin.GET("/users/:username", userController.GetUserByUsername)
	admin.GET("/users/:email", userController.GetUserByEmail)
	admin.PUT("/users/:id", userController.UpdateUser)

	// PRODUCT
	productController := controller.NewProductController(productUseCase)
	product := cfg.Echo.Group("/products", authMiddleware.IsAuthenticated())
	product.POST("", productController.CreateProductController, authMiddleware.IsAdmin)

	// MENU
	menuController := controller.NewMenuController(menuUseCase)
	menu := cfg.Echo.Group("/menus", authMiddleware.IsAuthenticated())
	menu.POST("", menuController.CreateMenuController, authMiddleware.IsAdmin)
	menu.GET("/:id", menuController.GetMenuByIDController)
	menu.GET("", menuController.GetAllMenusController)
	menu.GET("/:name", menuController.GetMenuByNameController)
	menu.GET("/:category", menuController.GetMenuByCategoryController)
	menu.PUT("/:id", menuController.UpdateMenuController, authMiddleware.IsAdmin)
	menu.DELETE("/:id", menuController.DeleteMenuByIDController, authMiddleware.IsAdmin)

}
