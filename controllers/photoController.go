package controllers

import (
	"net/http"
	"rakamin-submission/app"
	"rakamin-submission/database"
	"rakamin-submission/models"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	var postPhoto app.CreatePhoto

	err := c.ShouldBindJSON(&postPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	var photo = models.Photo{
		Title:    postPhoto.Title,
		Caption:  postPhoto.Caption,
		PhotoUrl: postPhoto.PhotoUrl,
	}
	err = database.DB.Create(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success upload photo",
	})
}
