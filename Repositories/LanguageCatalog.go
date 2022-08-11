package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetLanguageCatalogs(c *gin.Context) {
	languagecatalogs := []Models.LanguageCatalog{}
	Config.DB.Find(&languagecatalogs)
	c.JSON(200, &languagecatalogs)
}

func GetLanguageCatalog(c *gin.Context) {
	var languagecatalog Models.LanguageCatalog
	Config.DB.First(&languagecatalog, c.Param("id"))
	c.JSON(200, &languagecatalog)
}

func NewLanguageCatalog(c *gin.Context) {
	var languagecatalog Models.LanguageCatalog
	c.BindJSON(&languagecatalog)
	Config.DB.Create(&languagecatalog)
	c.JSON(200, &languagecatalog)
}

func UpdateLanguageCatalog(c *gin.Context) {
	var languagecatalog Models.LanguageCatalog
	Config.DB.Where("id = ?", c.Param("id")).First(&languagecatalog)
	c.BindJSON(&languagecatalog)
	Config.DB.Save(&languagecatalog)
	c.JSON(200, &languagecatalog)
}

func DeleteLanguageCatalog(c *gin.Context) {
	var languagecatalog Models.LanguageCatalog
	Config.DB.Where("id = ?", c.Param("id")).First(&languagecatalog)
	c.BindJSON(&languagecatalog)
	languagecatalog.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&languagecatalog)
	c.JSON(200, &languagecatalog)
}
