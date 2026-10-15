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
// @Description Просмотр атрибутов склада по его ID
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/ [get]
func (h *Handler) getWarehouseDetail(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
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
// @Description Создание роли пользователю по ID склада и ID пользователя
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param input body request.WarehouseEmployee true "Данные пользователя"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/employee/ [post]
func (h *Handler) createWarehouseEmployee(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}

	var input req.WarehouseEmployee

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.WarehouseID = warehouseID

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

type getWarehouseEmployeesResponse struct {
	Data []*models.WarehouseEmployee `json:"data"`
}

// @Summary Просмотр работников склада
// @Tags warehouse
// @Description Просмотр всех работников по ID склада
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/employee/ [get]
func (h *Handler) getWarehouseEmployees(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}

	warehouseEmployees, err := h.services.GetWarehouseEpmloyees(warehouseID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWarehouseEmployeesResponse{
		Data: warehouseEmployees,
	})
}

type getWarehouseEmployeeResponse struct {
	Data *models.WarehouseEmployee `json:"data"`
}

// @Summary Просмотр сотрудника склада
// @Tags warehouse
// @Description Просмотр сотрудника склада по ID склада и ID сотрудника
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param employeeID  path string true "Employee ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/employee/{employeeID} [get]
func (h *Handler) getWarehouseEmployee(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}
	employeeID, err := uuid.Parse(c.Param("employeeID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, "invalid employee id")
		return
	}

	employee, err := h.services.GetWarehouseEmployee(warehouseID, employeeID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWarehouseEmployeeResponse{
		Data: employee,
	})

}

// @Summary Просмотр сотрудника склада
// @Tags warehouse
// @Description Просмотр сотрудника склада по ID склада и ID сотрудника
// @Accept json
// @Produce json
// @Param warehouseID path string true "Warehouse ID"
// @Param employeeID  path string true "Employee ID"
// @Security ApiKeyAuth
// @Router /api/v1/warehouse/{warehouseID}/employee/{employeeID} [delete]
func (h *Handler) deleteWarehouseEmployee(c *gin.Context) {
	warehouseID, err := uuid.Parse(c.Param("warehouseID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid warehouse id")
		return
	}
	employeeID, err := uuid.Parse(c.Param("employeeID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, "invalid employee id")
		return
	}

	err = h.services.DeleteWarehouseEmployee(warehouseID, employeeID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusNoContent)

}
