package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
)

type getWarehouseDetailResponse struct {
	Data *models.Warehouse `json:"data"`
}

// @Summary Просмотр склада
// @Tags warehouse
// @Description Просмотр атрибутов склада по его id
// @Accept json
// @Produce json
// @Param id path string true "Warehouse ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/by-id/{id} [get]
func (h *Handler) getWarehouseDetail(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, "invalid warehouse id")
		return
	}

	warehouse, err := h.services.GetWarehouseDetail(&warehouseID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWarehouseDetailResponse{
		Data: warehouse,
	})

}

type createWarehouseEmployeeResponse struct {
	Data map[string]uuid.UUID `json:"data"`
}

// @Summary Создание роли пользователю
// @Tags warehouse
// @Description Создание роли пользователю по id склада и id пользователя
// @Accept json
// @Produce json
// @Param input body request.WarehouseEmployee true "Данные пользователя"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/employee [post]
func (h *Handler) createWarehouseEmployee(c *gin.Context) {
	var input req.WarehouseEmployee

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	warehouseEmployeeID, err := h.services.CreateWarhouseEmployee(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	data := map[string]uuid.UUID{
		"warehouseEmployeeID": warehouseEmployeeID,
	}
	c.JSON(http.StatusCreated, createWarehouseEmployeeResponse{
		Data: data,
	})

}
