package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIresponse("Register account failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIresponse("Register account failed",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest, response)
		return 
	}

	formatter := user.FormatterUser(newUser,"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

	response := helper.APIresponse("Account has been registered",http.StatusOK,"success",formatter)
	
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIresponse("Login failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil{
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIresponse("Login failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	formatter := user.FormatterUser(loggedinUser,"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

	response := helper.APIresponse("Successfuly loggedin",http.StatusOK,"success",formatter)
	
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context){
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIresponse("email checking failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	IsEmailAvalable, err := h.userService.IsEmailAvalable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIresponse("email checking failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	data := gin.H{
		"is_available" : IsEmailAvalable,
	}

	metaMessage := "Email has been registered"

	if IsEmailAvalable{
		metaMessage = "Email is available"
	}

	response := helper.APIresponse(metaMessage,http.StatusOK,"error",data)
	c.JSON(http.StatusOK, response)
	
}

func (h *userHandler) UploadAvatar(c *gin.Context){
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded":false}
		response := helper.APIresponse("Failed to upload avatar image", http.StatusBadRequest,"error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 31

	// path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-avatar-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded":false}
		response := helper.APIresponse("Failed to upload avatar image", http.StatusBadRequest,"error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded":false}
		response := helper.APIresponse("Failed to upload avatar image", http.StatusBadRequest,"error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded":true}
	response := helper.APIresponse("Avatar successfuly uploaded", http.StatusOK,"success", data)

	c.JSON(http.StatusOK, response)

}