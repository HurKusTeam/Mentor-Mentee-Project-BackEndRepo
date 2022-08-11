package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetMeetings(c *gin.Context) {
	meetings := []Models.Meeting{}
	Config.DB.Find(&meetings)
	c.JSON(200, &meetings)
}

func GetMeeting(c *gin.Context) {
	var meeting Models.Meeting
	Config.DB.First(&meeting, c.Param("id"))
	c.JSON(200, &meeting)
}

func NewMeeting(c *gin.Context) {
	var meeting Models.Meeting
	c.BindJSON(&meeting)
	Config.DB.Create(&meeting)
	c.JSON(200, &meeting)
}

func UpdateMeeting(c *gin.Context) {
	var meeting Models.Meeting
	Config.DB.Where("id = ?", c.Param("id")).First(&meeting)
	c.BindJSON(&meeting)
	Config.DB.Save(&meeting)
	c.JSON(200, &meeting)
}

func DeleteMeeting(c *gin.Context) {
	var meeting Models.Meeting
	Config.DB.Where("id = ?", c.Param("id")).First(&meeting)
	c.BindJSON(&meeting)
	meeting.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&meeting)
	c.JSON(200, &meeting)
}
