package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/sirupsen/logrus"
)

type createProductCategoryResponse struct {
	Data map[string]uuid.UUID `json:"data"`
}

// @Summary Создание категории продукта
// @Tags product
// @Description Создание категории продукта по ID склада
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param input body request.ProductCategory true "Данные категории"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/category/ [post]
func (h *Handler) createProductCategory(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}

	var input req.ProductCategory
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.WarehouseID = warehouseID

	categoryID, err := h.services.CreateProductCategory(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data := map[string]uuid.UUID{
		"categoryID": categoryID,
	}

	c.JSON(http.StatusCreated, createProductCategoryResponse{
		Data: data,
	})
}

type getProductCategoriesResponse struct {
	Data []*models.ProductCategory `json:"data"`
}

// @Summary Просмотр категорий продуктов
// @Tags product
// @Description Просмотр всех категорий по ID склада
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/category/ [get]
func (h *Handler) getProductCategories(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}

	categories, err := h.services.GetProductCategories(warehouseID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getProductCategoriesResponse{
		Data: categories,
	})
}

type getProductCategoryResponse struct {
	Data *models.ProductCategory `json:"data"`
}

// @Summary Просмотр категории продукта
// @Tags product
// @Description Просмотр категории по ID категории
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param categoryID path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/category/{categoryID}/ [get]
func (h *Handler) getProductCategory(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}

	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id")
		return
	}

	category, err := h.services.GetProductCategory(warehouseID, categoryID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getProductCategoryResponse{
		Data: category,
	})
}

// @Summary Удаление категории продукта
// @Tags product
// @Description Удаление категории по ID категории
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param categoryID path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/categorie/{categoryID}/ [delete]
func (h *Handler) deleteProductCategory(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}
	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id")
		return
	}

	if err := h.services.DeleteProductCategory(warehouseID, categoryID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

type createProductResponse struct {
	Data map[string]uuid.UUID `json:"data"`
}

// @Summary Создание продукта
// @Tags product
// @Description Создание продукта по ID категории продукта
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param categoryID path string true "Category ID"
// @Param input body request.Product true "Данные продукта"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/category/{categoryID}/product/ [post]
func (h *Handler) createProduct(c *gin.Context) {
	_, err := uuid.Parse(c.Param("warehouseID")) // TODO
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}
	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id")
		return
	}
	logrus.Info(categoryID)

	var input req.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.CategoryID = categoryID

	productID, err := h.services.CreateProduct(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data := map[string]uuid.UUID{
		"productID": productID,
	}

	c.JSON(http.StatusCreated, createProductResponse{
		Data: data,
	})
}

type getProductsResponse struct {
	Data []*models.Product `json:"data"`
}

// @Summary Просмотр продуктов склада по категории
// @Tags product
// @Description Просмотр всех продуктов по ID категории
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param categoryID path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/category/{categoryID}/product/ [get]
func (h *Handler) getProducts(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}
	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id")
		return
	}

	categories, err := h.services.GetProductsByCategory(warehouseID, categoryID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getProductsResponse{
		Data: categories,
	})
}

type getProductResponse struct {
	Data *models.Product `json:"data"`
}

// @Summary Просмотр продукта склада
// @Tags product
// @Description Просмотр продукта по ID склада, ID категории , ID продукта
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param categoryID path string true "Category ID"
// @Param productID path string true "Product ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/category/{categoryID}product/{productID} [get]
func (h *Handler) getProduct(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}
	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id")
		return
	}
	productID, err := uuid.Parse(c.Param("productID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id")
		return
	}

	product, err := h.services.GetProduct(warehouseID, categoryID, productID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getProductResponse{
		Data: product,
	})
}

// @Summary Удаление продукта
// @Tags product
// @Description Удаление продукта по ID склада, ID категории , ID продукта
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param categoryID path string true "Category ID"
// @Param productID path string true "Product ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/category/{categoryID}/product/{productID} [delete]
func (h *Handler) deleteProduct(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}
	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id")
		return
	}
	productID, err := uuid.Parse(c.Param("productID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id")
		return
	}

	if err := h.services.DeleteProduct(warehouseID, categoryID, productID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
