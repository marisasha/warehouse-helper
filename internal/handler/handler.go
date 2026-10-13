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
		wareHouse.GET("/:warehouseID", h.getWarehouseDetail)
		wareHouseEmployee := wareHouse.Group("/:warehouseID/employee")
		{
			wareHouseEmployee.POST("/", h.getWarehouseEmployees)
			wareHouseEmployee.GET("/", h.getWarehouseEmployees)
			wareHouseEmployee.GET("/:employeeID", h.getWarehouseEmployee)
			wareHouseEmployee.DELETE("/:employeeID", h.deleteWarehouseEmployee)
		}
		productCategorie := wareHouse.Group("/:warehouseID/categorie")
		{
			productCategorie.POST("/", h.createProductCategory)
			productCategorie.GET("/", h.getProductCategories)
			productCategorie.GET("/:categorieID", h.getProductCategory)
			productCategorie.DELETE("/:categorieID", h.deleteProductCategory)
		}
		product := wareHouse.Group("/:warehouseID/product")
		{
			product.POST("/")
			product.GET("/")
			product.GET("/:productID")
			product.DELETE("/:productID")
			productBatch := product.Group("/:productID/batch")
			{
				productBatch.POST("/")
				productBatch.GET("/")
				productBatch.GET("/:batchID")
			}
		}
	}
	supplier := apiV1.Group("/supplier")
	{
		supplier.GET("/:supplierID")
	}
	supply := apiV1.Group("/supply")
	{
		supply.POST("/")
		supply.GET("/")
		supply.GET("/:supplyID")
	}
	order := apiV1.Group("/order")
	{
		order.POST("/")
		order.GET("/")
		order.GET("/:orderID")
	}

	return router
}
