package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
)

type createSupplierResponse struct {
	Data map[string]uuid.UUID `json:"data"`
}

// @Summary Создание поставщика
// @Tags supply
// @Description Создание поставщика по id склада
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param input body request.Supplier true "Данные поставщика"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/supplier/ [post]
func (h *Handler) createSupplier(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "error in handler: invalid warehouse id")
		return
	}
	var input req.Supplier
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.WarehouseID = warehouseID

	supplierID, err := h.services.CreateSupplier(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data := map[string]uuid.UUID{
		"supplierID": supplierID,
	}
	c.JSON(http.StatusAccepted, createSupplierResponse{
		Data: data,
	})
}

type getSuppliersResponse struct {
	Data []*models.Supplier `json:"data"`
}

// @Summary Просмотр поставщиков склада
// @Tags supply
// @Description Просмотр всех работников по ID склада
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/supplier/ [get]
func (h *Handler) getSuppliers(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "error in handler: invalid warehouse id")
		return
	}

	suppliers, err := h.services.GetSuppliers(warehouseID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getSuppliersResponse{
		Data: suppliers,
	})
}

type getSupplierResponse struct {
	Data *models.Supplier `json:"data"`
}

// @Summary Просмотр поставщика склада
// @Tags supply
// @Description Просмотр поставщика склада по ID склада и ID поставщика
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param supplierID  path string true "Supplier ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/supplier/{supplierID}/ [get]
func (h *Handler) getSupplier(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "error in handler: nvalid warehouse id")
		return
	}
	supplierID, err := uuid.Parse(c.Param("supplierID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, "error in handler: invalid supplier id")
		return
	}

	supplier, err := h.services.GetSupplier(warehouseID, supplierID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getSupplierResponse{
		Data: supplier,
	})

}

// @Summary Удаление поставщика склада
// @Tags supply
// @Description Удаление поставщика склада по ID склада и ID поставщика
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param supplierID  path string true "Supplier ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/supplier/{supplierID}/ [delete]
func (h *Handler) deleteSupplier(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "error in handler: invalid warehouse id")
		return
	}
	supplierID, err := uuid.Parse(c.Param("supplierID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, "error in handler: invalid supplier id")
		return
	}

	err = h.services.DeleteSupplier(warehouseID, supplierID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusNoContent)

}

type createSupplyResponse struct {
	Data map[string]uuid.UUID `json:"data"`
}

// @Summary Создание поставки для склада
// @Tags supply
// @Description Создание поставка по id склада и id поставщика
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param supplierID path string true "Supplier ID"
// @Param input body request.Supply true "Данные для поставки"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/supplier/{supplierID}/supply/ [post]
func (h *Handler) createSupply(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "error in handler: invalid warehouse id")
		return
	}
	supplierID, err := uuid.Parse(c.Param("supplierID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "error in handler: invalid supplier id")
		return
	}
	var input req.Supply
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.SupplierID = supplierID
	input.CreatorID = userId

	supplyID, err := h.services.CreateSupply(&input, warehouseID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data := map[string]uuid.UUID{
		"supplyID": supplyID,
	}
	c.JSON(http.StatusAccepted, createSupplyResponse{
		Data: data,
	})
}
