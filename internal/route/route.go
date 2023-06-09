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
	categoryRepository := repository.NewCategoryRepository(cfg.DBConn)
	menuRepository := repository.NewMenuRepository(cfg.DBConn)

	userUseCase := usecase.NewUserUseCase(userRepository, cfg.JwtProvider)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
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

	// ADMIN
	admin := cfg.Echo.Group("/admin", authMiddleware.IsAdmin)
	admin.GET("/users", userController.GetAllUser)
	admin.GET("/users/:id", userController.GetUserByID)
	admin.GET("/users/:username", userController.GetUserByUsername)
	admin.GET("/users/:email", userController.GetUserByEmail)
	admin.PUT("/users/:id", userController.UpdateUser)

	// MENU

	// CATEGORY
	categoryController := controller.NewCategoryController(categoryUseCase)
	categoryRoutes := cfg.Echo.Group("/category")
	categoryRoutes.POST("", categoryController.CreateCategoryController, authMiddleware.IsAdmin)
	categoryRoutes.GET("", categoryController.GetCategoryController)
	categoryRoutes.GET("/id", categoryController.GetMenuByCategoryController)
	categoryRoutes.GET("/menu", categoryController.GetMenusByCategoryController)
	categoryRoutes.PUT("/:id", categoryController.UpdateCategoryController, authMiddleware.IsAdmin)
	categoryRoutes.DELETE("/:id", categoryController.DeleteCategoryController, authMiddleware.IsAdmin)

	// MENUS
	menuController := controller.NewMenuController(menuUseCase)
	menuRoutes := cfg.Echo.Group("/menus")
	menuRoutes.GET("", menuController.GetAllMenusController)                               //bisa
	menuRoutes.GET("/:id", menuController.GetMenuController)                               //bisa
	menuRoutes.GET("/name", menuController.GetMenusByNameController)                       //bisa
	menuRoutes.GET("/category", menuController.GetMenusByCategoryController)               //masih parsing error :V
	menuRoutes.GET("/category/name", menuController.GetMenusByCategoryNameController)      //bisa
	menuRoutes.POST("", menuController.CreateMenuController, authMiddleware.IsAdmin)       //bisa
	menuRoutes.PUT("/:id", menuController.UpdateMenuController, authMiddleware.IsAdmin)    //bisa
	menuRoutes.DELETE("/:id", menuController.DeleteMenuController, authMiddleware.IsAdmin) //bisa
}
