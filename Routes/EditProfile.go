package Routes

import (
	"TREgitim/Controllers"

	"github.com/gin-gonic/gin"
)

func EditProfileRoute(router *gin.Engine) {
	router.POST("/Profile", Controllers.EditProfile)
	router.GET("/Profile", Controllers.GetProfile)
	router.GET("/Profile/:id", Controllers.GetProfileElse)
	router.POST("/ProfileImg/", Controllers.ImageUpdate)
	router.POST("/UpdateSkill/", Controllers.UpdateSkill)
	router.POST("/UpdateLang/", Controllers.UpdateLang)
	router.POST("/AddSkill", Controllers.AddSkill)
}
