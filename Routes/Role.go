package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func RoleRoute(router *gin.Engine) {
	router.GET("/Role", Controllers.Role)
}
