package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	res "github.com/marisasha/warehouse-helper/internal/dto/request"
)

// @Summary Регистрация пользователя
// @Tags auth
// @Description Создание нового пользователя
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body request.User true "Данные пользователя"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input res.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.CreateUser(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]string{
		"message": "account sucsuccessfully created",
	})

}

// @Summary Аутентификация пользователя
// @Tags auth
// @Description Проверка прав пользователя
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body request.UserSignIn true "Данные пользователя"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input res.UserSignIn

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
