package api

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"

	"goshop/app/middleware"
)

//	@title			Blueprint Swagger API
//	@version		1.0
//	@description	Swagger API for Golang Project Blueprint.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	quangdangfit@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@BasePath	/api/v1

func RegisterAPI(r *gin.Engine, container *dig.Container) {
	err := container.Invoke(func(
		user *UserAPI,
		product *ProductAPI,
		order *OrderAPI,
	) {
		r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		authMiddleware := middleware.JWTAuth()
		refreshAuthMiddleware := middleware.JWTRefresh()
		authRoute := r.Group("/auth")
		{
			authRoute.POST("/register", user.Register)
			authRoute.POST("/login", user.Login)
			authRoute.POST("/refresh", refreshAuthMiddleware, user.RefreshToken)
			authRoute.GET("/me", authMiddleware, user.GetMe)
			authRoute.PUT("/change-password", authMiddleware, user.ChangePassword)
		}

		//--------------------------------API-----------------------------------
		api1 := r.Group("/api/v1")

		// Products
		productAPI := api1.Group("/products")
		{
			productAPI.GET("", product.ListProducts)
			productAPI.POST("", authMiddleware, product.CreateProduct)
			productAPI.PUT("/:id", authMiddleware, product.UpdateProduct)
			productAPI.GET("/:id", product.GetProductByID)
		}

		// Orders
		orderAPI := api1.Group("/orders", authMiddleware)
		{
			orderAPI.POST("", order.PlaceOrder)
			orderAPI.GET("/:id", order.GetOrderByID)
			orderAPI.GET("", order.GetOrders)
			orderAPI.PUT("/:id/cancel", order.CancelOrder)
		}
	})

	if err != nil {
		logger.Fatal(err)
	}
}

func Inject(container *dig.Container) {
	_ = container.Provide(NewProductAPI)
	_ = container.Provide(NewUserAPI)
	_ = container.Provide(NewOrderAPI)
}
