package controllers

import (
	"net/http"
	"rakamin-submission/app"
	"rakamin-submission/database"
	"rakamin-submission/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user app.UserRegis
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	userModel := models.User{
		Name:     user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	err = database.DB.Create(&userModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func GetAllUser(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	var userUp app.UserUpdate

	if err := c.ShouldBindJSON(&userUp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	var user = models.User{
		Name:     userUp.Username,
		Email:    userUp.Email,
		Password: userUp.Password,
	}

	update := database.DB.Model(&user).Where("id = ?", userId).Updates(&user)
	if update.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed Update User",
		})
		return
	}
	c.Set("user", nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "success update user",
	})

}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("userId")
	user := models.User{}

	if database.DB.Delete(&user, idParam).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed Delete User",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete user",
	})
}
