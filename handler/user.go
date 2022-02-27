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
	newData, err := h.userService.RegisterUserForm(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to create account", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(newData)
	response := helper.ApiResponse("Successfully created account", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
	return

}
func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Make sure you have filled out the form provided", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newData, err := h.userService.LoginUserForm(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Please check again your username or password", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(newData)
	response := helper.ApiResponse("Login Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return

}
func (h *userHandler) UpdateProfile(c *gin.Context) {
	var inputID user.DetailUserInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update profile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	var inputData user.UpdateInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update profile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newData, err := h.userService.UpdateUserForm(inputID, inputData)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to update profile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(newData)
	response := helper.ApiResponse("Update Profile Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
