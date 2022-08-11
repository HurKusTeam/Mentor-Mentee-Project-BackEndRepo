package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"TREgitim/Repositories"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisteringgUser struct {
	Username string
	Mail     string
	Password string
}

func AddMyMentor(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var user Models.User
	var reguser RegisteringgUser
	c.BindJSON(&reguser)

	Config.DB.First(&user, "mail=?", reguser.Mail)
	if user.Mail == reguser.Mail {
		c.JSON(200, "Bu mail adresi kullanılıyor.")
	}
	if user.Mail != reguser.Mail {
		password := fmt.Sprintf("%x", sha256.Sum256([]byte(reguser.Password)))
		user.Password = password
		user.UserName = reguser.Username
		user.Mail = reguser.Mail
		user.RegisterDate = time.Now()
		Repositories.NewUser(user)
	}
	var userr Models.User
	Config.DB.First(&userr, "mail=?", reguser.Mail)
	if userr.Mail != "" {
		var mentor Models.Mentor
		var comp Models.Company
		var usercomp Models.User
		Config.DB.Where("mail = ?", controll).First(&usercomp)
		Config.DB.Where("user_id = ?", usercomp.ID).First(&comp)
		mentor.UserID = userr.ID
		mentor.CompanyID = comp.ID
		Config.DB.Create(&mentor)
		c.JSON(202, "Kayıt başarılı")
	}
}

func MyMentors(c *gin.Context) {

	type Mentormodel struct {
		User    Models.User
		Mentor  Models.Mentor
		About   Models.About
		Profile Models.UserProfile
		Skills  []Models.SkillCatalog
	}
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)
	var user Models.User
	Config.DB.Where("mail = ?", controll).First(&user)

	var company Models.Company
	Config.DB.Where("user_id = ?", user.ID).First(&company)

	var mentors []Models.Mentor
	Config.DB.Where("company_id = ?", company.ID).Find(&mentors)

	var jsonmodel []Mentormodel
	for index, _ := range mentors {
		var usermodel Models.User
		var abouts Models.About
		var profiles Models.UserProfile
		var jsonmodell Mentormodel
		Config.DB.Where("id= ?", mentors[index].UserID).Find(&usermodel)
		Config.DB.Model(usermodel).Association("Abouts").Find(&abouts) // ilişki about
		Config.DB.Model(usermodel).Association("UserProfiles").Find(&profiles)
		//Config.DB.Model(usermodel).Association("Skills").Find(&skillsfirst)

		var skillsfirst []Models.Skill
		Config.DB.Where("user_id= ?", usermodel.ID).Find(&skillsfirst)
		var skillsfinal []Models.SkillCatalog
		skillsfinal = nil
		for index, _ := range skillsfirst {

			var skillssecond Models.SkillCatalog
			fmt.Println(skillsfirst[index].ID)
			Config.DB.Where("id= ?", skillsfirst[index].SkillCatalogID).First(&skillssecond)
			fmt.Println(skillssecond.Name)
			skillsfinal = append(skillsfinal, skillssecond)

		}

		jsonmodell.User = usermodel
		jsonmodell.About = abouts
		jsonmodell.Profile = profiles
		jsonmodell.Mentor = mentors[index]
		jsonmodell.Skills = skillsfinal

		jsonmodel = append(jsonmodel, jsonmodell)
	}

	c.JSON(202, jsonmodel)
}

func MyEmployees(c *gin.Context) {
	type jsonmodel struct {
		User Models.User
		Type string
	}

	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)
	if control == nil {
		c.Redirect(301, "/Login")
	}
	var user Models.User
	Config.DB.Where("mail = ?", controll).First(&user)

	var company Models.Company
	Config.DB.Where("user_id = ?", user.ID).First(&company)

	var mentors []Models.Mentor
	Config.DB.Where("company_id = ?", company.ID).Find(&mentors)
	var mentees []Models.Mentee
	Config.DB.Where("company_id = ?", company.ID).Find(&mentees)

	var userJson []jsonmodel
	for _, element := range mentors {
		var userj Models.User
		var userm jsonmodel //
		Config.DB.Where("id = ?", element.UserID).First(&userj)
		userm.User = userj
		userm.Type = "Mentor" //
		userJson = append(userJson, userm)
	}
	for _, element := range mentees {
		var userj Models.User
		var userm jsonmodel
		Config.DB.Where("id = ?", element.UserID).First(&userj)
		userm.User = userj
		userm.Type = "Mentee"
		userJson = append(userJson, userm)
	}

	//c.SetCookie("Mail", controll, 10, "", "", false, false)
	c.JSON(200, userJson)
}

func MatchMenteeMentor(c *gin.Context) {
	//id : mentorun userid
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)

	var mentor Models.Mentor
	var userskill Models.User
	number, _ := strconv.ParseUint(c.Param("mentorid"), 10, 32)
	Config.DB.Where("user_id = ?", uint(number)).First(&mentor)
	Config.DB.Where("id = ?", uint(number)).First(&userskill)
	var skills []Models.Skill
	Config.DB.Model(userskill).Association("Skills").Find(&skills)
	//fmt.Println(skills[0].SkillCatalogID)

	var user Models.User
	Config.DB.Where("mail = ?", controll).First(&user)

	var company Models.Company
	Config.DB.Where("user_id = ?", user.ID).First(&company)

	var mentees []Models.Mentee
	Config.DB.Where("company_id = ?", company.ID).Find(&mentees)

	type json struct {
		User        Models.UserProfile
		Menteeid    uint
		Trues       int
		Percent     int
		Mentorid    uint
		About       Models.About
		Userdefault Models.User
		Skills      []Models.SkillCatalog
		Mentee      Models.Mentee
	}

	//var menteestruecount []int
	var menteestruecount []json

	for _, element := range mentees {
		var userforeach Models.User // şirketin menteeleri mentenin user hali
		var skillsforeach []Models.Skill
		//var menteestruecount []int

		Config.DB.Where("id = ?", element.UserID).First(&userforeach)
		Config.DB.Model(userforeach).Association("Skills").Find(&skillsforeach) // menteenin yetenekler
		var trues int = 0
		for _, menteeskill := range skillsforeach { // menteeler donuyor
			//var trues int = 0
			for _, mentorskill := range skills { // mentorun yetenekler donuyor

				if menteeskill.SkillCatalogID == mentorskill.SkillCatalogID {
					trues += 1
				}
			}
			//menteestruecount = append(menteestruecount, trues)
		}

		var model json
		model.Menteeid = element.UserID
		model.Trues = trues

		//menteestruecount = append(menteestruecount, trues)
		menteestruecount = append(menteestruecount, model)
	}
	//fmt.Println(menteestruecount[1].Trues)
	for index, element := range menteestruecount {

		var user Models.UserProfile
		var x int
		x = 100 / len(skills)
		menteestruecount[index].Percent = element.Trues * x

		Config.DB.Where("user_id = ?", element.Menteeid).First(&user)
		menteestruecount[index].User = user
		menteestruecount[index].Mentorid = mentor.ID
		//
		var mentee Models.Mentee
		Config.DB.Where("user_id = ?", element.Menteeid).First(&mentee)
		menteestruecount[index].Mentee = mentee
		////

		var mentorr Models.Mentor
		Config.DB.Where("id = ?", mentor.ID).First(&mentorr)
		//fmt.Println(mentorr.Major)
		if mentee.Department == mentorr.Major && menteestruecount[index].Percent < 85 {
			menteestruecount[index].Percent += 10
		}
		if mentee.Department != mentorr.Major {
			menteestruecount[index].Percent -= 10
		}

		//
		var userr Models.User
		Config.DB.Where("id = ?", user.UserID).First(&userr)
		var abouts Models.About
		Config.DB.Model(userr).Association("Abouts").Find(&abouts)

		menteestruecount[index].About = abouts
		menteestruecount[index].Userdefault = userr

		var skillsfirst []Models.Skill
		var skillsfinal []Models.SkillCatalog
		Config.DB.Where("user_id= ?", userr.ID).Find(&skillsfirst)
		for index, _ := range skillsfirst {
			var skillssecond Models.SkillCatalog
			Config.DB.Where("id= ?", skillsfirst[index].SkillCatalogID).First(&skillssecond)
			skillsfinal = append(skillsfinal, skillssecond)

		}

		menteestruecount[index].Skills = skillsfinal

	}

	c.JSON(200, menteestruecount)
}

func ConfirmMenteeMentor(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	number, _ := strconv.ParseUint(c.Param("mentorid"), 10, 32)

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var mentee Models.Mentee
	Config.DB.Where("user_id = ?", c.Param("menteeid")).First(&mentee)
	mentee.MentorID = uint(number)
	Config.DB.Save(&mentee)

	c.JSON(200, "Mentor'a mentee başarıyla atandı.")
}

func MyMentorsPublic(c *gin.Context) {
	//user id company
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var comp Models.Company
	Config.DB.Where("user_id = ?", number).First(&comp)
	var mentors []Models.Mentor
	Config.DB.Where("company_id = ?", comp.ID).Find(&mentors)

	c.JSON(200, mentors)
}

func AddChildMentee(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)

	var user Models.User
	var reguser RegisteringgUser
	c.BindJSON(&reguser)

	Config.DB.First(&user, "mail=?", reguser.Mail)
	if user.Mail == reguser.Mail {
		c.JSON(200, "Bu mail adresi kullanılıyor.")
	}
	if user.Mail != reguser.Mail {
		password := fmt.Sprintf("%x", sha256.Sum256([]byte(reguser.Password)))
		user.Password = password
		user.UserName = reguser.Username
		user.Mail = reguser.Mail
		user.RegisterDate = time.Now()
		Repositories.NewUser(user)
	}

	var userr Models.User
	Config.DB.First(&userr, "mail=?", reguser.Mail)
	Config.DB.First("mail = ?", reguser.Mail).First(&userr)
	if userr.Mail != "" {
		var childment Models.Childmentee
		var mentee Models.Mentee
		var userrr Models.User
		Config.DB.Where("mail  = ? ", controll).First(&userrr)
		Config.DB.Where("user_id = ?", userrr.ID).First(&mentee)
		childment.UserID = userr.ID
		childment.MenteeID = mentee.ID
		Config.DB.Create(&childment)
		c.JSON(202, "Kayıt başarılı")
	}

}
