package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func GetAllSkillsRoute(router *gin.Engine) {

	router.GET("/GetSkills", Controllers.GetAllSkills)
}
