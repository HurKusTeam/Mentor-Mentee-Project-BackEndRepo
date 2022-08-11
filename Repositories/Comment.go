package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	comments := []Models.Comment{}
	Config.DB.Find(&comments)
	c.JSON(200, &comments)
}

func GetComment(c *gin.Context) {
	var comment Models.Comment
	Config.DB.First(&comment, c.Param("id"))
	c.JSON(200, &comment)
}

func NewComment(c *gin.Context) {
	var comment Models.Comment
	c.BindJSON(&comment)
	Config.DB.Create(&comment)
	c.JSON(200, &comment)
}

func UpdateComment(c *gin.Context) {
	var comment Models.Comment
	Config.DB.Where("id = ?", c.Param("id")).First(&comment)
	c.BindJSON(&comment)
	Config.DB.Save(&comment)
	c.JSON(200, &comment)
}

func DeleteComment(c *gin.Context) {
	var comment Models.Comment
	Config.DB.Where("id = ?", c.Param("id")).First(&comment)
	c.BindJSON(&comment)
	comment.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&comment)
	c.JSON(200, &comment)
}
