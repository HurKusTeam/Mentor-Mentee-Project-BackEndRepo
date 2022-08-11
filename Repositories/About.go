package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetAbouts(c *gin.Context) {
	abouts := []Models.About{}
	Config.DB.Find(&abouts)
	c.JSON(200, &abouts)
}

func GetAbout(c *gin.Context) {
	var about Models.About
	Config.DB.First(&about, c.Param("id"))
	c.JSON(200, &about)
}

func NewAbout(c *gin.Context) {
	var about Models.About
	c.BindJSON(&about)
	Config.DB.Create(&about)
	c.JSON(200, &about)
}

func UpdateAbout(c *gin.Context) {
	var about Models.About
	Config.DB.Where("id = ?", c.Param("id")).First(&about)
	c.BindJSON(&about)
	Config.DB.Save(&about)
	c.JSON(200, &about)
}

func DeleteAbout(c *gin.Context) {
	var about Models.About
	Config.DB.Where("id = ?", c.Param("id")).First(&about)
	c.BindJSON(&about)
	about.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&about)
	c.JSON(200, &about)
}
