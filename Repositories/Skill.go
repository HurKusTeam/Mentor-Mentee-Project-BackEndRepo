package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetSkills(c *gin.Context) {
	skills := []Models.Skill{}
	Config.DB.Find(&skills)
	c.JSON(200, &skills)
}

func GetSkill(c *gin.Context) {
	var skill Models.Skill
	Config.DB.First(&skill, c.Param("id"))
	c.JSON(200, &skill)
}

func NewSkill(c *gin.Context) {
	var skill Models.Skill
	c.BindJSON(&skill)
	Config.DB.Create(&skill)
	c.JSON(200, &skill)
}

func UpdateSkill(c *gin.Context) {
	var skill Models.Skill
	Config.DB.Where("id = ?", c.Param("id")).First(&skill)
	c.BindJSON(&skill)
	Config.DB.Save(&skill)
	c.JSON(200, &skill)
}

func DeleteSkill(c *gin.Context) {
	var skill Models.Skill
	Config.DB.Where("id = ?", c.Param("id")).First(&skill)
	c.BindJSON(&skill)
	skill.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&skill)
	c.JSON(200, &skill)
}
