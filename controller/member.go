package controller

import (
	"church-ws/models"
	"church-ws/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /members
// Get all user
func FindMembers(c *gin.Context) {
	var users []models.Member
	models.DB.Find(&users)
	if len(users) > 0 {
		c.JSON(http.StatusOK, util.GenerateResponse(0, "success", users))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(1, "no data found", nil))
}

// GET /member/:id
// Get member by id
func FindMember(c *gin.Context) {
	var member models.Member
	if err := models.DB.Where("id = ?", c.Param("id")).First(&member).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "member not found", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", member))
}

// POST /member
// Create a member
func CreateMember(c *gin.Context) {
	var input models.CreateMember
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	dob, err := util.ParseDate(input.Dob)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	now := time.Now()
	member := models.Member{
		Surname:         input.Surname,
		Othernames:      input.Othernames,
		Dob:             &dob,
		Gender:          input.Gender,
		Maritalstatus:   input.MaritalStatus,
		Employed:        input.Employed,
		Occupation:      input.Occupation,
		Company:         input.Company,
		Companylocation: input.CompanyLocation,
		Residence:       input.Residence,
		Mobile:          input.Mobile,
		Email:           input.Email,
		Datecreated:     &now,
		Presbytery:      input.Presbytery,
		Status:          "ACTIVE",
		Firstname:       input.Firstname,
	}
	if result := models.DB.Create(&member); result.Error != nil {
		fmt.Printf("failed to create member: %+v\n", result.Error.Error())
		c.JSON(http.StatusOK, util.GenerateResponse(1, "failed to create member", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}

// PUT /member
// Update member
func UpdateMember(c *gin.Context) {
	var input models.CreateMember
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	dob, err := util.ParseDate(input.Dob)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	fmt.Printf("updating %+v\n", input)
	//Get model if exist
	var member models.Member
	if err := models.DB.Where("id = ?", input.Id).First(&member).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "member not found", nil))
		return
	}
	member.Surname = input.Surname
	member.Firstname = input.Firstname
	member.Othernames = input.Othernames
	member.Dob = &dob
	member.Gender = input.Gender
	member.Maritalstatus = input.MaritalStatus
	member.Employed = input.Employed
	member.Occupation = input.Occupation
	member.Company = input.Company
	member.Companylocation = input.CompanyLocation
	member.Residence = input.Residence
	member.Mobile = input.Mobile
	member.Email = input.Email
	member.Presbytery = input.Presbytery
	now := time.Now()
	member.ModifiedDate = &now
	if result := models.DB.Save(&member); result.Error != nil {
		fmt.Printf("failed to update member: %+v\n", result.Error.Error())
		c.JSON(http.StatusOK, util.GenerateResponse(1, "failed to update member", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}

// DELETE /member/:id
// Delete a member
func DeleteMember(c *gin.Context) {
	var member models.Member
	if err := models.DB.Where("id = ?", c.Param("id")).First(&member).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "member not found", nil))
		return
	}
	models.DB.Delete(&member)
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}
