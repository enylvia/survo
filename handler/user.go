package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"survorest/helper"
	"survorest/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create account", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.ApiResponse("Success to create account", http.StatusOK, "success", input)
	c.JSON(http.StatusOK, response)

}
func (h *userHandler) LoginUser(c *gin.Context) {
	//panic("not implemented yet")
}
func (h *userHandler) UpdateProfile(c *gin.Context) {
	//panic("not implemented yet")
}
