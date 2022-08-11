package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MenteeModel struct {
	Name        string
	Surname     string
	Department  string
	University  string
	Mail        string
	PhoneNumber string
	GPA         float64
	City        string
	Linkedin    string
	Github      string
	MenteeID    uint
	Skill       []string
	ProfilPhoto string
	UserID      uint
}

func MenteeList(c *gin.Context) {
	var mentmodels []MenteeModel

	var mentees []Models.Mentee
	var mentor Models.Mentor
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionid"]
	fmt.Println(control)
	Config.DB.Where("user_id = ?", control).First(&mentor)
	Config.DB.Where("mentor_id = ?", mentor.ID).Find(&mentees)
	for _, mentvar := range mentees {
		var mentmodel MenteeModel
		var uni Models.University
		var unic Models.UniversityCatalog
		var userp Models.UserProfile
		var about Models.About
		var skillnames []string
		var u Models.User

		var skills []Models.Skill
		Config.DB.Where("user_id = ?", mentvar.UserID).First(&uni)
		Config.DB.Where("id = ?", uni.UniversityCatalogID).First(&unic)
		Config.DB.Where("user_id = ?", mentvar.UserID).First(&userp)
		Config.DB.Where("user_id = ?", mentvar.UserID).First(&about)
		Config.DB.Where("user_id = ?", mentvar.UserID).Find(&skills)
		Config.DB.Where("id = ?", mentvar.UserID).First(&u)

		for _, sk := range skills {
			var skillc Models.SkillCatalog
			Config.DB.Where("id = ?", sk.SkillCatalogID).First(&skillc)
			skillnames = append(skillnames, skillc.Name)
		}

		mentmodel.Name = userp.Name
		mentmodel.Surname = userp.Surname
		mentmodel.City = userp.City
		mentmodel.Department = mentvar.Department
		mentmodel.GPA = mentvar.GPA
		mentmodel.University = unic.Name
		mentmodel.ProfilPhoto = userp.ProfileImage
		mentmodel.MenteeID = mentvar.ID
		mentmodel.Linkedin = about.Linkedin
		mentmodel.Github = about.GitHub
		mentmodel.Skill = skillnames
		mentmodel.PhoneNumber = userp.PhoneNumber
		mentmodel.Mail = u.Mail
		mentmodel.UserID = u.ID
		mentmodels = append(mentmodels, mentmodel)
	}

	c.JSON(202, mentmodels)
}
