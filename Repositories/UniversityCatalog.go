package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetUniversityCatalogs(c *gin.Context) {
	universitycatalogs := []Models.UniversityCatalog{}
	Config.DB.Find(&universitycatalogs)
	c.JSON(200, &universitycatalogs)
}

func GetUniversityCatalog(c *gin.Context) {
	var universitycatalog Models.UniversityCatalog
	Config.DB.First(&universitycatalog, c.Param("id"))
	c.JSON(200, &universitycatalog)
}

func NewUniversityCatalog(c *gin.Context) {
	var universitycatalog Models.UniversityCatalog
	c.BindJSON(&universitycatalog)
	Config.DB.Create(&universitycatalog)
	c.JSON(200, &universitycatalog)
}

func UpdateUniversityCatalog(c *gin.Context) {
	var universitycatalog Models.UniversityCatalog
	Config.DB.Where("id = ?", c.Param("id")).First(&universitycatalog)
	c.BindJSON(&universitycatalog)
	Config.DB.Save(&universitycatalog)
	c.JSON(200, &universitycatalog)
}

func DeleteUniversityCatalog(c *gin.Context) {
	var universitycatalog Models.UniversityCatalog
	Config.DB.Where("id = ?", c.Param("id")).First(&universitycatalog)
	c.BindJSON(&universitycatalog)
	//universitycatalog.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&universitycatalog)
	c.JSON(200, &universitycatalog)
}
