package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func IndividualMentorRoute(router *gin.Engine) {

	router.GET("/IndividualMentors", Controllers.GetIndividual)
}
