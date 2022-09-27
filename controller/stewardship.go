package controller

import (
	"church-ws/models"
	"church-ws/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /stewardgroups
// Get all steward groups
func FindStewardGroups(c *gin.Context) {
	var groups []models.Stewardgroup
	models.DB.Find(&groups)
	if len(groups) > 0 {
		c.JSON(http.StatusOK, util.GenerateResponse(0, "success", groups))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(1, "no data found", nil))
}

// GET /stewardgroup/:id
// Get a steward group
func FindStewardGroup(c *gin.Context) {
	var group models.Stewardgroup
	if err := models.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "group not found", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", group))
}

// POST /stewardgroup
// Create steward group
func CreateStewardGroup(c *gin.Context) {
	var input models.CreateStewardGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	var leader models.Member
	if err := models.DB.Where("id = ?", input.LeaderId).First(&leader).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "leader not found", nil))
		return
	}
	now := time.Now()
	group := models.Stewardgroup{
		Name:        input.Name,
		Leader:      leader.Id,
		LeaderName:  fmt.Sprintf("%s %s", leader.Firstname, leader.Surname),
		DateCreated: &now,
		Status:      "ACTIVE",
	}
	if result := models.DB.Create(&group); result.Error != nil {
		fmt.Printf("failed to create group: %+v\n", result.Error.Error())
		c.JSON(http.StatusOK, util.GenerateResponse(1, "failed to create group", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}

// PUT /user/:id
// Update a user
func UpdateStewardGroup(c *gin.Context) {
	//Validate input
	var input models.CreateStewardGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(1, "invalid request body", nil))
		return
	}
	fmt.Printf("updating %+v\n", input)
	var leader models.Member
	if err := models.DB.Where("id = ?", input.LeaderId).First(&leader).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "leader not found", nil))
		return
	}
	//Get model if exist
	var group models.Stewardgroup
	if err := models.DB.Where("id = ?", input.Id).First(&group).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "group not found", nil))
		return
	}
	now := time.Now()
	group.Leader = leader.Id
	group.LeaderName = fmt.Sprintf("%s %s", leader.Firstname, leader.Surname)
	group.ModifiedDate = &now
	group.Name = input.Name
	if result := models.DB.Save(&group); result.Error != nil {
		fmt.Printf("failed to update group: %+v\n", result.Error.Error())
		c.JSON(http.StatusOK, util.GenerateResponse(1, "failed to update group", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}

// DELETE /stewardgroup/:id
// Delete a stewardgroup
func DeleteStewardGroup(c *gin.Context) {
	var group models.Stewardgroup
	if err := models.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusOK, util.GenerateResponse(1, "group not found", nil))
		return
	}
	models.DB.Delete(&group)
	c.JSON(http.StatusOK, util.GenerateResponse(0, "success", nil))
}
