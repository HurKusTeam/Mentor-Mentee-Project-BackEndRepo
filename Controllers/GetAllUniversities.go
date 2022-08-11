package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"

	"github.com/gin-gonic/gin"
)

func GetAllUniversities(c *gin.Context) {
	var unic []Models.UniversityCatalog
	Config.DB.Find(&unic)
	c.JSON(200, &unic)
}
