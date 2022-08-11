package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetSkillCatalogs(c *gin.Context) {
	skillcatalogs := []Models.SkillCatalog{}
	Config.DB.Find(&skillcatalogs)
	c.JSON(200, &skillcatalogs)
}

func GetSkillCatalog(c *gin.Context) {
	var skillcatalog Models.SkillCatalog
	Config.DB.First(&skillcatalog, c.Param("id"))
	c.JSON(200, &skillcatalog)
}

func NewSkillCatalog(c *gin.Context) {
	var skillcatalog Models.SkillCatalog
	c.BindJSON(&skillcatalog)
	Config.DB.Create(&skillcatalog)
	c.JSON(200, &skillcatalog)
}

func UpdateSkillCatalog(c *gin.Context) {
	var skillcatalog Models.SkillCatalog
	Config.DB.Where("id = ?", c.Param("id")).First(&skillcatalog)
	c.BindJSON(&skillcatalog)
	Config.DB.Save(&skillcatalog)
	c.JSON(200, &skillcatalog)
}

func DeleteSkillCatalog(c *gin.Context) {
	var skillcatalog Models.SkillCatalog
	Config.DB.Where("id = ?", c.Param("id")).First(&skillcatalog)
	c.BindJSON(&skillcatalog)
	//skillcatalog.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&skillcatalog)
	c.JSON(200, &skillcatalog)
}
