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
	menuItemRepository := repository.NewMenuItemRepository(cfg.DBConn)
	orderItemRepository := repository.NewOrderItemRepository(cfg.DBConn)
	orderRepository := repository.NewOrderRepository(cfg.DBConn)
	paymentRepository := repository.NewPaymentRepository(cfg.DBConn)

	userUseCase := usecase.NewUserUseCase(userRepository, cfg.JwtProvider)
	menuItemUseCase := usecase.NewMenuItemUseCase(menuItemRepository)
	orderItemUseCase := usecase.NewOrderItemUseCase(orderItemRepository)
	orderUseCase := usecase.NewOrderUseCase(orderRepository)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository)

	// Routes

	// AUTH
	authController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/register", authController.RegisterUserController)
	cfg.Echo.POST("/login", authController.LoginController)
	// AUTH ADMIN
	adminController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/admin/register", adminController.RegisterAdminController)

	// USER
	userController := controller.NewUserController(userUseCase)
	user := cfg.Echo.Group("/users", authMiddleware.IsAuthenticated())
	user.GET("", userController.GetAllUser, authMiddleware.IsAdmin)
	user.GET("/:id", userController.GetUserByID, authMiddleware.IsUser)

	// MENU ITEM
	menuItemController := controller.NewMenuItemController(menuItemUseCase)
	menu := cfg.Echo.Group("/menu-items", authMiddleware.IsAuthenticated())
	menu.GET("", menuItemController.GetAllMenuItems)
	menu.GET("/:id", menuItemController.GetMenuItemByID)
	menu.POST("/", menuItemController.CreateMenuItem, authMiddleware.IsAdmin)
	menu.PUT("/:id", menuItemController.UpdateMenuItem, authMiddleware.IsAdmin)
	menu.DELETE("/:id", menuItemController.DeleteMenuItem, authMiddleware.IsAdmin)

	// ORDER ITEM
	orderItemController := controller.NewOrderItemController(orderItemUseCase)
	orderItem := cfg.Echo.Group("/order-items", authMiddleware.IsAuthenticated())
	orderItem.GET("/:id", orderItemController.GetOrderItemsByOrderID)
	orderItem.POST("/", orderItemController.AddOrderItem, authMiddleware.IsAdmin)
	orderItem.PUT("/:id", orderItemController.UpdateOrderItem, authMiddleware.IsAdmin)
	orderItem.DELETE("/:id", orderItemController.RemoveOrderItem)

	// ORDER
	orderController := controller.NewOrderController(orderUseCase)
	order := cfg.Echo.Group("/orders", authMiddleware.IsAuthenticated())
	order.GET("/:id", orderController.GetOrderByID)
	order.POST("/", orderController.CreateOrder)

	// PAYMENT
	paymentController := controller.NewPaymentController(paymentUseCase)
	payments := cfg.Echo.Group("/payments", authMiddleware.IsAuthenticated())
	payments.GET("", paymentController.GetAllPayments, authMiddleware.IsAdmin)
	payments.GET("/:id", paymentController.GetPaymentByID)
	payments.POST("/", paymentController.CreatePayment)
	payments.PUT("/:id", paymentController.UpdatePayment)
	payments.DELETE("/:id", paymentController.DeletePayment, authMiddleware.IsAdmin)
}
