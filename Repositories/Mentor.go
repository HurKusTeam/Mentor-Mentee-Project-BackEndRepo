package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetMentors(c *gin.Context) {
	mentors := []Models.Mentor{}
	Config.DB.Find(&mentors)
	c.JSON(200, &mentors)
}

func GetMentor(c *gin.Context) Models.Mentor {
	var mentor Models.Mentor
	Config.DB.First(&mentor, c.Param("id"))
	//c.JSON(200, &mentor)
	return mentor
}

func NewMentor(c *gin.Context) {
	var mentor Models.Mentor
	c.BindJSON(&mentor)
	Config.DB.Create(&mentor)
	c.JSON(200, &mentor)
}

func UpdateMentor(c *gin.Context) {
	var mentor Models.Mentor
	Config.DB.Where("id = ?", c.Param("id")).First(&mentor)
	c.BindJSON(&mentor)
	Config.DB.Save(&mentor)
	c.JSON(200, &mentor)
}

func Deletementor(c *gin.Context) {
	var mentor Models.Mentor
	Config.DB.Where("id = ?", c.Param("id")).First(&mentor)
	c.BindJSON(&mentor)
	mentor.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&mentor)
	c.JSON(200, &mentor)
}
