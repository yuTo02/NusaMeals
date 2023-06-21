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
	paymentRepository := repository.NewPaymentRepository(cfg.DBConn)
	orderRepository := repository.NewOrderRepository(cfg.DBConn)
	cartRepository := repository.NewCartRepository(cfg.DBConn)

	userUseCase := usecase.NewUserUseCase(userRepository, cfg.JwtProvider)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository, menuRepository)
	menuUseCase := usecase.NewMenuUseCase(menuRepository)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, menuRepository)
	cartUseCase := usecase.NewCartUseCase(cartRepository, menuRepository)

	// Routes

	// AUTH
	authController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/login", authController.LoginController)
	cfg.Echo.POST("/logout", authController.LogoutController)
	//register ADMIN
	adminController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/admin/register", adminController.RegisterAdminController)
	//register USER
	cfg.Echo.POST("/register", authController.RegisterUserController)

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

	// CATEGORY
	categoryController := controller.NewCategoryController(categoryUseCase)
	categoryRoutes := cfg.Echo.Group("/category")
	categoryRoutes.POST("", categoryController.CreateCategoryController, authMiddleware.IsAdmin)
	categoryRoutes.GET("", categoryController.GetCategoriesController)
	categoryRoutes.GET("/id", categoryController.GetMenusByCategoryIDController)
	categoryRoutes.GET("/menu", categoryController.GetMenusByCategoryNameController)
	categoryRoutes.PUT("/:id", categoryController.UpdateCategoryController, authMiddleware.IsAdmin)
	categoryRoutes.DELETE("/:id", categoryController.DeleteCategoryController, authMiddleware.IsAdmin)

	// MENUS
	menuController := controller.NewMenuController(menuUseCase)
	menuRoutes := cfg.Echo.Group("/menus")
	menuRoutes.GET("", menuController.GetAllMenusController)                               //bisa
	menuRoutes.GET("/:id", menuController.GetMenuController)                               //bisa
	menuRoutes.GET("/name", menuController.GetMenusByNameController)                       //bisa
	menuRoutes.GET("/category", menuController.GetMenusByCategoryController)               //bisa
	menuRoutes.GET("/category/name", menuController.GetMenusByCategoryNameController)      //bisa
	menuRoutes.POST("", menuController.CreateMenuController, authMiddleware.IsAdmin)       //bisa
	menuRoutes.PUT("/:id", menuController.UpdateMenuController, authMiddleware.IsAdmin)    //bisa
	menuRoutes.DELETE("/:id", menuController.DeleteMenuController, authMiddleware.IsAdmin) //bisa

	//PAYMENT
	paymentController := controller.NewPaymentController(paymentUseCase)
	paymentRoutes := cfg.Echo.Group("/payments", authMiddleware.IsAuthenticated())
	paymentRoutes.POST("", paymentController.CreatePayment)
	paymentRoutes.PUT("/:id", paymentController.UpdatePayment)
	paymentRoutes.PUT("/admin/:id", paymentController.UpdatePaymentByAdmin, authMiddleware.IsAdmin)
	paymentRoutes.DELETE("/:id", paymentController.DeletePayment)
	paymentRoutes.GET("/:id", paymentController.GetPaymentByID)
	paymentRoutes.GET("", paymentController.GetAllPayments)

	//ORDER
	orderController := controller.NewOrderController(orderUseCase)
	orderRoutes := cfg.Echo.Group("/orders", authMiddleware.IsAuthenticated())
	orderRoutes.POST("", orderController.CreateOrder)
	orderRoutes.GET("/:orderID", orderController.GetOrderByID)
	orderRoutes.PUT("/:orderID", orderController.UpdateOrder, authMiddleware.IsAdmin)
	orderRoutes.DELETE("/:orderID", orderController.DeleteOrder, authMiddleware.IsAdmin)
	orderRoutes.GET("", orderController.GetAllOrders, authMiddleware.IsAdmin)

	// CART
	cartController := controller.NewCartController(cartUseCase)
	cartRoutes := cfg.Echo.Group("/carts", authMiddleware.IsAuthenticated())
	cartRoutes.GET("/:cartID/total", cartController.GetCartTotal)
	cartRoutes.POST("/items", cartController.AddItemToCart)
	cartRoutes.PUT("/items/:cartID", cartController.UpdateCartItemQuantity)
	cartRoutes.GET("/:cartID/items", cartController.GetCartItems)
	cartRoutes.DELETE("/items/:cartID", cartController.RemoveItemFromCart)
	cartRoutes.DELETE("/:cartID", cartController.ClearCart)

}
