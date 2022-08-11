package Routes

import (
	"TREgitim/Controllers"
	"github.com/gin-gonic/gin"
)

func Company(router *gin.Engine) {
	router.GET("/MyMentors", Controllers.MyMentors)
	router.GET("/MatchMenteeMentor/:mentorid", Controllers.MatchMenteeMentor)
	router.GET("/MyEmployees", Controllers.MyEmployees)
	router.GET("/ConfirmMenteeMentor/:menteeid/:mentorid", Controllers.ConfirmMenteeMentor)
	router.POST("/AddMyMentor", Controllers.AddMyMentor)
	router.GET("/MyMentorsPublic/:id", Controllers.MyMentorsPublic)
	router.POST("/AddChildMentee", Controllers.AddChildMentee)
}
