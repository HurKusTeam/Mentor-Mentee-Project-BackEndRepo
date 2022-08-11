package Routes

import (
	"TREgitim/Controllers"
	"github.com/gin-gonic/gin"
)

func Application(router *gin.Engine) {
	router.GET("/MakeAnApplication/:id", Controllers.MakeAnApplication)
	router.GET("/MyApplications/:id", Controllers.MyApplications)
	router.GET("/ReceivedApplications/:id", Controllers.ReceivedApplications)
	router.GET("/AcceptApplication/:id/:uid/:tid", Controllers.AcceptApplication)
	router.GET("/RejectApplication/:id/:uid", Controllers.RejectApplication)
}
