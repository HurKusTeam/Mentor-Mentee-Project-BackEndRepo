package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func MentorMenteeListRoute(router *gin.Engine) {

	router.GET("/MenteeList", Controllers.MenteeList)
}
