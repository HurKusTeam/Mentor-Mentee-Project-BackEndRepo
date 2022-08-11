package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"

	"github.com/gin-gonic/gin"
)

func GetAllSkills(c *gin.Context) {
	var skillc []Models.SkillCatalog
	Config.DB.Find(&skillc)
	c.JSON(200, &skillc)
}
