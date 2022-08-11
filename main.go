package main

import (
	"TREgitim/Config"
	"TREgitim/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	Config.Connect()
	Routes.LoginRoute(router)
	Routes.RegisterRoute(router)
	Routes.AdvertRoute(router)
	Routes.Application(router)
	Routes.Company(router)
	Routes.Todo(router)
	Routes.MentorMenteeListRoute(router)
	Routes.EditProfileRoute(router)
	Routes.MeetingRoute(router)
	Routes.DeleteRoute(router)
	Routes.IndividualMentorRoute(router)
	Routes.GetAllLanguagesRoute(router)
	Routes.GetAllSkillsRoute(router)
	Routes.GetAllUniversitiesRoute(router)
	router.Run(":8080")
}
