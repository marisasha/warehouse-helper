package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/sirupsen/logrus"
)

type getWarehouseDetailResponse struct {
	Data *models.Warehouse
}

// @Summary Просмотр склада
// @Tags warehouse
// @Description Просмотр атрибутов склада по его id
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{id} [get]
func (h *Handler) getWarehouseDetail(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	logrus.Info(warehouseID)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	warehouse, err := h.services.GetWarehouseDetail(&warehouseID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, getWarehouseDetailResponse{
		Data: warehouse,
	})

}
