package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
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
// @Router /api/v1/warehouse/{warehouseID}/categorie [post]
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
// @Description Просмотр всех категорий по id склада
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/categorie [get]
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
// @Description Просмотр категории по id склада и id категории
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param categoryID path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/categorie/{categoryID} [get]
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
// @Description Удаление категории по id склада и id категории
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param categoryID path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/categorie/{categoryID} [delete]
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
