package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"

	"github.com/gin-gonic/gin"
)

func GetAllLanguages(c *gin.Context) {
	var langc []Models.LanguageCatalog
	Config.DB.Find(&langc)
	c.JSON(200, langc)
}
