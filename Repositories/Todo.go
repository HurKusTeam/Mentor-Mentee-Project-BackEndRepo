package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos := []Models.Todo{}
	Config.DB.Find(&todos)
	c.JSON(200, &todos)
}

func GetTodo(c *gin.Context) {
	var todo Models.Todo
	Config.DB.First(&todo, c.Param("id"))
	c.JSON(200, &todo)
}

func NewTodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)
	Config.DB.Create(&todo)
	c.JSON(200, &todo)
}

func UpdateTodo(c *gin.Context) {
	var todo Models.Todo
	Config.DB.Where("id = ?", c.Param("id")).First(&todo)
	c.BindJSON(&todo)
	Config.DB.Save(&todo)
	c.JSON(200, &todo)
}

func DeleteTodo(c *gin.Context) {
	var todo Models.Todo
	Config.DB.Where("id = ?", c.Param("id")).First(&todo)
	c.BindJSON(&todo)
	todo.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&todo)
	c.JSON(200, &todo)
}
