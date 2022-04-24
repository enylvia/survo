package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"survorest/auth"
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
	formatter := user.FormatUserRegister(newData)
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
	token := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}
	tokenString, err := token.GenerateToken(int(newData.Id), newData.Email)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to generate token", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(newData, tokenString)
	response := helper.ApiResponse("Login Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return

}
func (h *userHandler) UpdateProfile(c *gin.Context) {
	var inputID user.DetailUserInput
	currentUser := c.MustGet("claims").(user.User)

	var userId int
	userId = int(currentUser.Id)

	err := c.ShouldBindUri(&inputID)
	if userId != inputID.ID {
		errorMessage := gin.H{"error": "You are not authorized to update this user"}
		response := helper.ApiResponse("You are not authorized to update this user", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	var inputData user.UpdateInput
	err = c.ShouldBindJSON(&inputData)
	newData, err := h.userService.UpdateUserForm(inputID, inputData)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to update profile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatDetailUser(newData)
	response := helper.ApiResponse("Update Profile Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
func (h *userHandler) UploadAvatar(c *gin.Context) {
	var inputID user.DetailUserInput
	currentUser := c.MustGet("claims").(user.User)

	var userId int
	userId = int(currentUser.Id)

	err := c.ShouldBindUri(&inputID)

	if userId != inputID.ID {
		errorMessage := gin.H{"error": "You are not authorized to update this user"}
		response := helper.ApiResponse("You are not authorized to update this user", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update profile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	file, err := c.FormFile("image")
	fileName := strings.Join(strings.Fields(file.Filename), "")
	pathFile := fmt.Sprintf("images/%d-%s", userId, fileName)
	err = c.SaveUploadedFile(file, pathFile)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to save uploaded image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newData, err := h.userService.UploadAvatar(inputID.ID, pathFile)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to upload avatar", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatDetailUser(newData)
	response := helper.ApiResponse("Update Profile Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
func (h *userHandler) GetProfile(c *gin.Context) {
	var inputID user.DetailUserInput
	currentUser := c.MustGet("claims").(user.User)

	var userId int
	userId = int(currentUser.Id)

	err := c.ShouldBindUri(&inputID)

	if userId != inputID.ID {
		errorMessage := gin.H{"error": "You are not authorized to see this user profile"}
		response := helper.ApiResponse("You are not authorized to see this user profile", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("ID Invalid", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	getData, err := h.userService.GetUserByID(inputID.ID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("User with that ID not found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatDetailUser(getData)
	response := helper.ApiResponse("Successfully get data", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
