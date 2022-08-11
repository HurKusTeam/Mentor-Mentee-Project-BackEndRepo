package Routes

import (
	"TREgitim/Controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {
	router.POST("/Register", Controllers.Register)
}
