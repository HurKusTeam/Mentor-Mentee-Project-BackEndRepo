package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetAdverts(c *gin.Context) []Models.Advert {
	adverts := []Models.Advert{}
	Config.DB.Find(&adverts)
	//c.JSON(200, &adverts)
	return adverts
}

func GetAdvert(c *gin.Context) Models.Advert {
	var advert Models.Advert
	Config.DB.First(&advert, c.Param("id"))
	//c.JSON(200, &advert)
	return advert
}

func NewAdvert(c *gin.Context) {
	var advert Models.Advert
	c.BindJSON(&advert)
	Config.DB.Create(&advert)
	c.JSON(200, &advert)
}

func UpdateAdvert(c *gin.Context) {
	var advert Models.Advert
	Config.DB.Where("id = ?", c.Param("id")).First(&advert)
	c.BindJSON(&advert)
	Config.DB.Save(&advert)
	c.JSON(200, &advert)
}

func DeleteAdvert(c *gin.Context) {
	var advert Models.Advert
	Config.DB.Where("id = ?", c.Param("id")).First(&advert)
	c.BindJSON(&advert)
	advert.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&advert)
	c.JSON(200, &advert)
}
