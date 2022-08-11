package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func MeetingRoute(router *gin.Engine) {
	router.POST("/Meetings/:id", Controllers.CreateMeeting)

	router.GET("/Meetings/:id", Controllers.GetMeetings)

}
