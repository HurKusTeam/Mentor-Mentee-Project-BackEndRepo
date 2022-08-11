package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetLanguages(c *gin.Context) {
	languages := []Models.Language{}
	Config.DB.Find(&languages)
	c.JSON(200, &languages)
}

func GetLanguage(c *gin.Context) {
	var language Models.Language
	Config.DB.First(&language, c.Param("id"))
	c.JSON(200, &language)
}

func NewLanguage(c *gin.Context) {
	var language Models.Language
	c.BindJSON(&language)
	Config.DB.Create(&language)
	c.JSON(200, &language)
}

func UpdateLanguage(c *gin.Context) {
	var language Models.Language
	Config.DB.Where("id = ?", c.Param("id")).First(&language)
	c.BindJSON(&language)
	Config.DB.Save(&language)
	c.JSON(200, &language)
}

func DeleteLanguage(c *gin.Context) {
	var language Models.Language
	Config.DB.Where("id = ?", c.Param("id")).First(&language)
	c.BindJSON(&language)
	language.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&language)
	c.JSON(200, &language)
}
