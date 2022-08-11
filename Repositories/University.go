package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetUniversitys(c *gin.Context) {
	universities := []Models.University{}
	Config.DB.Find(&universities)
	c.JSON(200, &universities)
}

func GetUniversity(c *gin.Context) {
	var university Models.University
	Config.DB.First(&university, c.Param("id"))
	c.JSON(200, &university)
}

func NewUniversity(c *gin.Context) {
	var university Models.University
	c.BindJSON(&university)
	Config.DB.Create(&university)
	c.JSON(200, &university)
}

func UpdateUniversity(c *gin.Context) {
	var university Models.University
	Config.DB.Where("id = ?", c.Param("id")).First(&university)
	c.BindJSON(&university)
	Config.DB.Save(&university)
	c.JSON(200, &university)
}

func DeleteUniversity(c *gin.Context) {
	var university Models.University
	Config.DB.Where("id = ?", c.Param("id")).First(&university)
	c.BindJSON(&university)
	university.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&university)
	c.JSON(200, &university)
}
