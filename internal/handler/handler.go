package handler

import (
	"github.com/marisasha/warehouse-helper/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.Use(h.loggingMiddleware)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	apiV1 := router.Group("/api/v1")
	{
		wareHouse := apiV1.Group("/warehouse")
		wareHouse.GET("/by-id/:id", h.getWarehouseDetail)
		wareHouseEmployee := wareHouse.Group("/employee")
		{
			wareHouseEmployee.POST("/", h.createWarehouseEmployee)
			wareHouseEmployee.GET("/")
			wareHouseEmployee.GET("/:id")
			wareHouseEmployee.DELETE("/:id")
		}
		productCategorie := wareHouse.Group("/categorie")
		{
			productCategorie.POST("/")
			productCategorie.GET("/")
			productCategorie.GET("/:id")
			productCategorie.DELETE("/:id")
		}
		product := wareHouse.Group("/product")
		{
			product.POST("/")
			product.GET("/")
			product.GET("/by-id/:id")
			product.DELETE("/by-id/:id")
			productBatch := product.Group("/batch")
			{
				productBatch.POST("/")
				productBatch.GET("/")
				productBatch.GET("/:id")
			}
		}
	}
	supplier := apiV1.Group("/supplier")
	{
		supplier.GET("/:id")
	}
	supply := apiV1.Group("/supply")
	{
		supply.POST("/")
		supply.GET("/")
		supply.GET("/:id")
	}
	order := apiV1.Group("/order")
	{
		order.POST("/")
		order.GET("/")
		order.GET("/:id")
	}

	return router
}
