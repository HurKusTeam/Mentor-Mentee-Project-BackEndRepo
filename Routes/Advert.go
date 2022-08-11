package Routes

import (
	"TREgitim/Controllers"
	"github.com/gin-gonic/gin"
)

func AdvertRoute(router *gin.Engine) {
	router.POST("/AddAdvert/:id", Controllers.AddAdvert)
	//router.GET("/GetAdvert/:id", Controllers.GetAdvert)
	//router.GET("/GetAllAdverts", Controllers.GetAllAdverts)
	router.GET("/GetAdvertSolo/:id", Controllers.GetAdvertSolo)
	router.GET("GetAdvertsAll", Controllers.GetAdvertsAll)

}
