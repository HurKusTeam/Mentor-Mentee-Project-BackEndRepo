package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"

	"github.com/gin-gonic/gin"
)

type WhatRole struct {
	Role int
}

func Role(c *gin.Context) {
	var mentor Models.Mentor
	var mentee Models.Mentee
	var company Models.Company
	var rl WhatRole

	rl.Role = 0
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionid"]

	Config.DB.First(&mentor, "user_id = ?", control)
	Config.DB.First(&mentee, "user_id = ?", control)
	Config.DB.First(&company, "user_id = ?", control)

	if mentor.ID != 0 {
		rl.Role = 1
	}
	if mentee.ID != 0 {
		rl.Role = 2
	}
	if company.ID != 0 {
		rl.Role = 3
	}

	c.ShouldBindJSON(&rl)
	c.JSON(200, rl)
}
