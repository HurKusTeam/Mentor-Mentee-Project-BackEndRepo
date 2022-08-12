package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"

	"github.com/gin-gonic/gin"
)

func GetMenteeListForMeeting(c *gin.Context) {
	var mentees []Models.Mentee
	var mentor Models.Mentor
	Config.DB.First(&mentor, c.Param("id"))
	Config.DB.Where("mentor_id = ?", c.Param("id")).Find(&mentees)
	Config.DB.Model(mentor).Association("Mentees").Find(&mentees)

	c.JSON(202, gin.H{
		"mentees": mentees,
	})

}

func GetMeetings(c *gin.Context) {

	var mentee Models.Mentee
	var mentor Models.Mentor
	var menteeCont Models.Mentee
	var meetings []Models.Meeting
	i := c.Param("id")
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionid"]
	Config.DB.First(&menteeCont, "user_id", control)
	Config.DB.First(&mentor, "user_id", control)

	if menteeCont.ID != 0 {
		Config.DB.Where("mentee_id = ? AND mentor_id = ?", menteeCont.ID, menteeCont.MentorID).Find(&meetings)

	}
	if mentor.ID != 0 {
		Config.DB.Where("id = ?", i).First(&mentee)
		Config.DB.Where("mentee_id = ? AND mentor_id = ?", mentee.ID, mentee.MentorID).Find(&meetings)
	}

	c.JSON(200, meetings)

}

func CreateMeeting(c *gin.Context) {
	var mentee Models.Mentee
	var meeting Models.Meeting
	var modelmeeting Models.Meeting

	Config.DB.First(&mentee, "id = ?", c.Param("id"))

	c.ShouldBindJSON(&modelmeeting)
	meeting.Title = modelmeeting.Title
	meeting.Link = modelmeeting.Link
	meeting.Date = modelmeeting.Date
	meeting.MenteeID = mentee.ID
	meeting.Description = modelmeeting.Description
	meeting.MentorID = mentee.MentorID

	Config.DB.Create(&meeting)
	c.ShouldBindJSON(&meeting)
	c.JSON(200, &meeting)

}
