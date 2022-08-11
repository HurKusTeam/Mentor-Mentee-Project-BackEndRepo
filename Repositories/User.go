package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []Models.User{}
	Config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetUser(c *gin.Context) Models.User {
	var user Models.User
	Config.DB.First(&user, c.Param("id"))
	//c.JSON(200, &user)
	return user
}

func NewUser(user Models.User) {
	//var user Models.User
	//c.BindJSON(&user)
	Config.DB.Create(&user)
	//c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user Models.User
	Config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	Config.DB.Save(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	Config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	user.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&user)
	c.JSON(200, &user)
}
