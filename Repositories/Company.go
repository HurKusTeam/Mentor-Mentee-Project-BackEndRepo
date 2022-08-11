package Repositories

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"github.com/gin-gonic/gin"
)

func GetCompanys(c *gin.Context) {
	companies := []Models.Company{}
	Config.DB.Find(&companies)
	c.JSON(200, &companies)
}

func GetCompany(c *gin.Context) Models.Company {
	var company Models.Company
	Config.DB.First(&company, c.Param("id"))
	//c.JSON(200, &company)
	return company
}

func NewCompany(c *gin.Context) {
	var company Models.Company
	c.BindJSON(&company)
	Config.DB.Create(&company)
	c.JSON(200, &company)
}

func UpdateCompany(c *gin.Context) {
	var company Models.Company
	Config.DB.Where("id = ?", c.Param("id")).First(&company)
	c.BindJSON(&company)
	Config.DB.Save(&company)
	c.JSON(200, &company)
}

func DeleteCompany(c *gin.Context) {
	var company Models.Company
	Config.DB.Where("id = ?", c.Param("id")).First(&company)
	c.BindJSON(&company)
	company.IsDeleted = true // hata çıkarsa burada
	Config.DB.Save(&company)
	c.JSON(200, &company)
}
