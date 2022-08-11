package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func DeleteRoute(router *gin.Engine) {

	router.DELETE("/DeleteSkills/:name", Controllers.DeleteSkill)
	router.DELETE("/DeleteLanguage/:name", Controllers.DeleteLanguage)

}
