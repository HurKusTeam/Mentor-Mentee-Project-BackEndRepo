package Routes

import (
	"TREgitim/Controllers"
	"github.com/gin-gonic/gin"
)

func Todo(router *gin.Engine) {
	router.GET("/MyMentees", Controllers.MyMentees)
	router.POST("/AddTodo", Controllers.AddTodo)
	router.GET("/GetTodo/:menteeid/:mentorid", Controllers.GetTodo)
	router.GET("/DeleteTodo/:id", Controllers.DeleteTodo)
	router.POST("/UpdateTodo/:id", Controllers.UpdateTodo)
	router.POST("/Dragged", Controllers.Dragged)
	router.POST("/AddComment", Controllers.AddComment)
	router.GET("/TodoComments/:id", Controllers.TodoComments)

}
