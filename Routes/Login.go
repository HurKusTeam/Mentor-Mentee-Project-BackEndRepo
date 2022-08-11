package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func LoginRoute(router *gin.Engine) {
	router.POST("/Login", Controllers.Login)
	router.GET("/Logout", Controllers.Logout)
	router.GET("/Logout/:id", Controllers.Logout)

}
