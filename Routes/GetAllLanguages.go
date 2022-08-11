package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func GetAllLanguagesRoute(router *gin.Engine) {

	router.GET("/GetLanguages/", Controllers.GetAllLanguages)
}
