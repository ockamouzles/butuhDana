package handler

import (
	"net/http"
	"projekStartup/helper"
	"projekStartup/user"

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
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("U have mistake", http.StatusUnprocessableEntity, "Error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.Register(input)

	if err != nil {
		err := helper.APIResponse("U have mistake", http.StatusBadRequest, "Failed", nil)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	formatter := user.FormatUser(newUser, "apapunamanyaaa")

	response := helper.APIResponse("Sukses didaftarkan", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusAccepted, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Erroors!!!", http.StatusUnprocessableEntity, "Error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("U mistake, u stupid idiot", http.StatusUnprocessableEntity, "Error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, "blablasldsjdasjdda")

	response := helper.APIResponse("Sucessfully logged", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// package handler

// import (
// 	"fmt"
// 	"net/http"
// 	"projekStartup/user"

// 	"github.com/gin-gonic/gin"
// )

// // buat struct(wadah) untuk struct dari service
// type userHandler struct {
// 	userService user.Service
// }

// // hubungkan userHandler dengan userService (user.Service)
// func NewUserHandler(userService user.Service) *userHandler {
// 	return &userHandler{userService}
// }

// // Processing dimulai
// func (h *userHandler) RegisterUser(c *gin.Context) {
// 	var input user.RegisterUserInput //tangkap data dari RegisterUserInput
// 	err := c.ShouldBindJSON(&input)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, nil)
// 	}

// 	user, err := h.userService.Register(input)
// 	fmt.Println(user)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, nil)
// 	}

// 	c.JSON(http.StatusOK, user)
// }
