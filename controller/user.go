package controller

import (
	"church-ws/models"
	"church-ws/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /users
// Get all user
func FindUsers(c *gin.Context) {
	fmt.Println("finding all user")
	var users []models.User
	models.DB.Find(&users)
	if len(users) > 0 {
		c.JSON(http.StatusOK, util.GenerateResponse(0, "success", users))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(1, "no data found", nil))
}

// GET /user/:id
// Find a book
func FindUser(c *gin.Context) {
	fmt.Println("finding a user")
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "user not found", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", user))
}

// POST /user
// Create a user
func CreateUser(c *gin.Context) {
	fmt.Println("creating user")
	var input models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	now := time.Now()
	user := models.User{
		Username:    input.Username,
		Password:    util.EncryptPassword(input.Password),
		CreatedBy:   1,
		Datecreated: &now,
		Msisdn:      input.Msisdn,
		Status:      "ACTIVE",
		Firstname:   input.Firstname,
		Lastname:    input.Lastname,
	}
	if result := models.DB.Create(&user); result.Error != nil {
		fmt.Printf("failed to create user: %+v\n", result.Error.Error())
		c.JSON(http.StatusOK, util.GenerateResponse(1, "failed to create user", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}

// PUT /user/:id
// Update a user
func UpdateUser(c *gin.Context) {
	fmt.Println("updating user")
	//Validate input
	var input models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	fmt.Printf("updating %+v\n", input)
	//Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", input.Id).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "user not found", nil))
		return
	}
	user.Firstname = input.Firstname
	user.Lastname = input.Lastname
	user.Msisdn = input.Msisdn
	user.Username = input.Username
	now := time.Now()
	user.Lastupdatetime = &now
	if result := models.DB.Save(&user); result.Error != nil {
		fmt.Printf("failed to update user: %+v\n", result.Error.Error())
		c.JSON(http.StatusOK, util.GenerateResponse(1, "failed to update user", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}

// DELETE /user/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	fmt.Println("deleting user")
	// db.Delete(&User{}, 10)
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "user not found", nil))
		return
	}
	models.DB.Delete(&user)
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}

// POST /user/login
// Login for user
func Login(c *gin.Context) {
	fmt.Println("login")
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	var user models.User
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "Invalid username or password", nil))
		return
	}
	verified, _ := util.VerifyPassword(input.Password, user.Password)
	if !verified {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "Invalid username or password", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}

// POST /user/changepasswd
func ChangePassword(c *gin.Context) {
	fmt.Println("change password")
	var input models.ChangePassword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	var user models.User
	if err := models.DB.Where("id = ?", input.Id).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "user not found", nil))
		return
	}
	verified, _ := util.VerifyPassword(input.Oldpassword, user.Password)
	if !verified {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "Password does much previous", nil))
		return
	}
	user.Password = util.EncryptPassword(input.Newpassword)
	now := time.Now()
	user.Lastupdatetime = &now
	if result := models.DB.Save(&user); result.Error != nil {
		fmt.Printf("failed to update user password: %+v\n", result.Error.Error())
		c.JSON(http.StatusOK, util.GenerateResponse(1, "failed to update user password", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}
