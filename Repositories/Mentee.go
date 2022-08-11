package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetMentees(c *gin.Context) {
	mentees := []Models.Mentee{}
	Config.DB.Find(&mentees)
	c.JSON(200, &mentees)
}

func GetMentee(c *gin.Context) {
	var mentee Models.Mentee
	Config.DB.First(&mentee, c.Param("id"))
	c.JSON(200, &mentee)
}

func NewMentee(c *gin.Context) {
	var mentee Models.Mentee
	c.BindJSON(&mentee)
	Config.DB.Create(&mentee)
	c.JSON(200, &mentee)
}

func UpdateMentee(c *gin.Context) {
	var mentee Models.Mentee
	Config.DB.Where("id = ?", c.Param("id")).First(&mentee)
	c.BindJSON(&mentee)
	Config.DB.Save(&mentee)
	c.JSON(200, &mentee)
}

func DeleteMentee(c *gin.Context) {
	var mentee Models.Mentee
	Config.DB.Where("id = ?", c.Param("id")).First(&mentee)
	c.BindJSON(&mentee)
	mentee.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&mentee)
	c.JSON(200, &mentee)
}
