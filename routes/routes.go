package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	controllers "api-finalproject/controller"
	"api-finalproject/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/product", controllers.GetAllProduct)
	r.GET("/product/:id", controllers.GetProductById)
	r.POST("/product", controllers.CreateProduct)
	r.PATCH("/product/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)

	r.GET("/product-categories", controllers.GetAllCategory)
	r.GET("/product-categories/:id", controllers.GetCategoryById)
	r.GET("/product-categories/:id/product", controllers.GetProductByCategoryId)
	r.POST("/product-categories", controllers.CreateCategory)
	r.PATCH("/product-categories/:id", controllers.UpdateCategory)
	r.DELETE("/product-categories/:id", controllers.DeleteCategory)

	r.GET("/orderline", controllers.GetAllOrderline)
	r.GET("/orderline/:id", controllers.GetOrderlineById)

	orderlineMiddlewareRoute := r.Group("/orderline")
	orderlineMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	orderlineMiddlewareRoute.POST("/", controllers.CreateOrderline)
	orderlineMiddlewareRoute.PATCH("/:id", controllers.UpdateOrderline)
	orderlineMiddlewareRoute.DELETE("/:id", controllers.DeleteOrderline)

	r.GET("/order", controllers.GetAllOrder)
	r.GET("/order/:id", controllers.GetOrderById)
	r.GET("/order/:id/orderline", controllers.GetOrderlineByOderId)

	orderMiddlewareRoute := r.Group("/order")
	orderMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	orderMiddlewareRoute.POST("/", controllers.CreateOrder)
	orderMiddlewareRoute.PATCH("/:id", controllers.UpdateOrder)
	orderMiddlewareRoute.DELETE("/:id", controllers.DeleteOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
