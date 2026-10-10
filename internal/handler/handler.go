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
		wareHouse.GET("/:wareHouseID", h.getWarehouseDetail)
		wareHouseEmployee := wareHouse.Group("/:wareHouseID/employee")
		{
			wareHouseEmployee.POST("/")
			wareHouseEmployee.GET("/")
			wareHouseEmployee.GET("/:ID")
			wareHouseEmployee.DELETE("/:ID")
		}
		productCategorie := wareHouse.Group("/:wareHouseID/categorie")
		{
			productCategorie.POST("/")
			productCategorie.GET("/")
			productCategorie.GET("/:ID")
			productCategorie.DELETE("/:ID")
		}
		product := wareHouse.Group("/:wareHouseID/product")
		{
			product.POST("/")
			product.GET("/")
			product.GET("/:productID")
			product.DELETE("/:productID")
			productBatch := product.Group("/:productID/batch")
			{
				productBatch.POST("/")
				productBatch.GET("/")
				productBatch.GET("/:ID")
			}
		}
	}
	supplier := apiV1.Group("/supplier")
	{
		supplier.GET("/:ID")
	}
	supply := apiV1.Group("/supply")
	{
		supply.POST("/")
		supply.GET("/")
		supply.GET("/:ID")
	}
	order := apiV1.Group("/order")
	{
		order.POST("/")
		order.GET("/")
		order.GET("/:ID")
	}

	return router
}
