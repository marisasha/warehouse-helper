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
			wareHouseEmployee.POST("/", h.createWarehouseEmployee)
			wareHouseEmployee.GET("/", h.getWarehouseEmployees)
			wareHouseEmployee.GET("/:employeeID", h.getWarehouseEmployee)
			wareHouseEmployee.DELETE("/:employeeID", h.deleteWarehouseEmployee)
		}
		productCategorie := wareHouse.Group("/:warehouseID/category")
		{
			productCategorie.POST("/", h.createProductCategory)
			productCategorie.GET("/", h.getProductCategories)
			productCategorie.GET("/:categoryID", h.getProductCategory)
			productCategorie.DELETE("/:categoryID", h.deleteProductCategory)
			product := productCategorie.Group("/:categoryID/product")
			{
				product.POST("/", h.createProduct)
				product.GET("/", h.getProducts)
				product.GET("/:productID", h.getProduct)
				product.DELETE("/:productID", h.deleteProduct)
				productBatch := product.Group("/:productID/batch")
				{
					productBatch.POST("/")
					productBatch.GET("/")
					productBatch.GET("/:batchID")
				}
			}
		}
		supplier := wareHouse.Group("/:warehouseID/supplier")
		{
			supplier.POST("/", h.createSupplier)
			supplier.GET("/", h.getSuppliers)
			supplier.GET("/:supplierID", h.getSupplier)
			supplier.DELETE("/:supplierID", h.deleteSupplier)
			supply := supplier.Group("/:supplierID/supply", h.userIdentity)
			{
				supply.POST("/", h.createSupply)
				supply.GET("/")
				supply.GET("/:supplyID")
			}
		}
	}
	order := apiV1.Group("/order")
	{
		order.POST("/")
		order.GET("/")
		order.GET("/:orderID")
	}

	return router
}
