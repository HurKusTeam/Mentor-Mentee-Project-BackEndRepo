package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func GetAllUniversitiesRoute(router *gin.Engine) {

	router.GET("/GetUniversities", Controllers.GetAllUniversities)
}
