package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"TREgitim/Repositories"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AppModel struct {
	User          Models.User
	Advertidmodel uint
	About         Models.About
	Profile       Models.UserProfile
	Skills        []Models.SkillCatalog
}

func MakeAnApplication(c *gin.Context) {
	var application Models.Application
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var user Models.User
	Config.DB.Where("mail = ?", controll).First(&user)
	fmt.Println(user.ID)
	fmt.Println(user.Mail)
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	application.AdvertID = uint(number)
	application.UserID = user.ID
	application.Status = 1
	application.Date = time.Now().Format("02-01-2006") // sabah denenecek
	Config.DB.Create(&application)
	c.JSON(200, &application)
	//c.Redirect(301, "/GetAllAdverts")

}

type AppModell struct {
	Applications Models.Application
	UserProf     Models.UserProfile
	Company      Models.Company
}

func MyApplications(c *gin.Context) {
	//session, _ := store.Get(c.Request, "sessioncontrol")
	//control := session.Values["sessionmail"]

	var applications = Repositories.GetApplications(c)
	//number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var jsonmodel []AppModell
	for _, element := range applications {
		var solomodel AppModell
		solomodel.Applications = element
		var advert Models.Advert
		Config.DB.Where("id = ?", element.AdvertID).First(&advert)
		var comp Models.Company
		Config.DB.Where("id = ?", advert.CompanyID).First(&comp)
		var prof Models.UserProfile
		Config.DB.Where("user_id = ?", comp.UserID).First(&prof)
		solomodel.UserProf = prof
		solomodel.Company = comp

		jsonmodel = append(jsonmodel, solomodel)
	}

	c.JSON(200, jsonmodel)
}

func ReceivedApplications(c *gin.Context) {
	var userr Models.User
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)
	Config.DB.Where("mail = ?", controll).First(&userr)
	var numb = userr.ID
	var num = uint(numb)

	var comp Models.Company
	var advert Models.Advert
	Config.DB.Where("user_id = ?", num).First(&comp)
	if comp.Description != "" {
		number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		Config.DB.Where("company_id = ?", uint(number)).First(&advert)
	}
	if comp.Description == "" {
		number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		Config.DB.Where("mentor_id = ?", uint(number)).First(&advert)
	}
	var applications []Models.Application
	var users []Models.User
	numberr := uint(advert.ID)
	Config.DB.Where("advert_id = ? AND status = ?", numberr, 1).Find(&applications)
	for _, element := range applications {
		var user Models.User
		numberrr := uint(element.UserID)
		Config.DB.Where("id = ?", (numberrr)).First(&user)
		users = append(users, user)
	}

	var listmodel []AppModel // sonrası
	for _, element := range applications {
		var user Models.User
		numberrr := uint(element.UserID)
		Config.DB.Where("id = ?", (numberrr)).First(&user)
		var solomodel AppModel
		var about Models.About
		var profiles Models.UserProfile
		Config.DB.Model(user).Association("Abouts").Find(&about) //
		Config.DB.Model(user).Association("UserProfiles").Find(&profiles)

		var skillsfirst []Models.Skill
		Config.DB.Where("user_id= ?", user.ID).Find(&skillsfirst)
		var skillsfinal []Models.SkillCatalog
		skillsfinal = nil
		for index, _ := range skillsfirst {

			var skillssecond Models.SkillCatalog
			fmt.Println(skillsfirst[index].ID)
			Config.DB.Where("id= ?", skillsfirst[index].SkillCatalogID).First(&skillssecond)
			fmt.Println(skillssecond.Name)
			skillsfinal = append(skillsfinal, skillssecond)

		}

		solomodel.About = about
		solomodel.Profile = profiles
		solomodel.User = user
		solomodel.Advertidmodel = numberr
		solomodel.Skills = skillsfinal

		listmodel = append(listmodel, solomodel) // sonrası

	}

	c.JSON(200, listmodel)
}

func AcceptApplication(c *gin.Context) {
	//advertid/userid/comp(userid)
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	numberrr, _ := strconv.ParseUint(c.Param("uid"), 10, 32)
	tid, _ := strconv.ParseUint(c.Param("tid"), 10, 32)
	var mentee Models.Mentee
	var application Models.Application
	var comp Models.Company
	var mentor Models.Mentor
	Config.DB.Where("user_id = ?", uint(tid)).First(&comp)
	Config.DB.Where("user_id = ?", uint(tid)).First(&mentor)

	if comp.Description != "" {
		mentee.CompanyID = comp.ID
		mentee.UserID = uint(numberrr)
		mentee.MentorID = 56 // sabit
		Config.DB.Create(&mentee)
	}

	if comp.Description == "" {
		mentee.MentorID = mentor.ID
		mentee.UserID = uint(numberrr)
		mentee.CompanyID = 3 // sabit
		Config.DB.Create(&mentee)
	}

	Config.DB.Where("advert_id = ? AND user_id = ? ", uint(number), uint(numberrr)).First(&application)
	application.Status = 2
	Config.DB.Save(&application)
	c.JSON(200, application)
	//c.Redirect(202, "/ReceivedApplications/11")

}

func RejectApplication(c *gin.Context) {

	var application Models.Application
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	numberrr, _ := strconv.ParseUint(c.Param("uid"), 10, 32)
	numberr := uint(number)
	numberrrr := uint(numberrr)
	Config.DB.Where("advert_id = ? AND user_id = ? ", numberr, numberrrr).First(&application)

	application.Status = 3
	Config.DB.Save(&application)
	c.JSON(200, application)
	//c.Redirect(202, "/ReceivedApplications/11")
}
