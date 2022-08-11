package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetApplications(c *gin.Context) []Models.Application {
	applications := []Models.Application{}
	Config.DB.Where("user_id = ?", c.Param("id")).Find(&applications)
	//Config.DB.Find(&applications)
	return applications
	//c.JSON(200, &applications)
}

func GetApplication(c *gin.Context) {
	var application Models.Application
	Config.DB.First(&application, c.Param("id"))
	c.JSON(200, &application)
}

func NewApplication(c *gin.Context) {
	var application Models.Application
	c.BindJSON(&application)
	Config.DB.Create(&application)
	c.JSON(200, &application)
}

func UpdateApplication(c *gin.Context) {
	var application Models.Application
	Config.DB.Where("id = ?", c.Param("id")).First(&application)
	c.BindJSON(&application)
	Config.DB.Save(&application)
	c.JSON(200, &application)
}

func DeleteApplication(c *gin.Context) {
	var application Models.Application
	Config.DB.Where("id = ?", c.Param("id")).First(&application)
	c.BindJSON(&application)
	application.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&application)
	c.JSON(200, &application)
}
