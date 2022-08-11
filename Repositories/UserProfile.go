package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetUserProfiles(c *gin.Context) {
	userprofiles := []Models.UserProfile{}
	Config.DB.Find(&userprofiles)
	c.JSON(200, &userprofiles)
}

func GetUserProfile(c *gin.Context) {
	var userprofile Models.UserProfile
	Config.DB.First(&userprofile, c.Param("id"))
	c.JSON(200, &userprofile)
}

func NewUserProfile(c *gin.Context) {
	var userprofile Models.UserProfile
	c.BindJSON(&userprofile)
	Config.DB.Create(&userprofile)
	c.JSON(200, &userprofile)
}

func UpdateUserProfile(c *gin.Context) {
	var userprofile Models.UserProfile
	Config.DB.Where("id = ?", c.Param("id")).First(&userprofile)
	c.BindJSON(&userprofile)
	Config.DB.Save(&userprofile)
	c.JSON(200, &userprofile)
}

func DeleteUserProfile(c *gin.Context) {
	var userprofile Models.UserProfile
	Config.DB.Where("id = ?", c.Param("id")).First(&userprofile)
	c.BindJSON(&userprofile)
	userprofile.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&userprofile)
	c.JSON(200, &userprofile)
}
