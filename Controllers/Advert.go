package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"TREgitim/Repositories"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdvertModel struct {
	Title       string
	Description string
	Startdate   string // time time dı ikiside
	Enddate     string
}

func AddAdvert(c *gin.Context) {
	// parametre olarak olarak company veya mentor tablosundaki userid gelecek id değil.
	//session, _ := store.Get(c.Request, "sessioncontrol")
	//control := session.Values["sessionmail"]
	var advertmodel AdvertModel
	var company Models.Company
	var mentor Models.Mentor
	var advert Models.Advert

	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	Config.DB.Where("user_id = ?", number).First(&company)
	//company = Repositories.GetCompany(c)

	if company.ID == 0 {
		Config.DB.Where("user_id = ?", number).First(&mentor)
		//mentor = Repositories.GetMentor(c)
		c.BindJSON(&advertmodel)
		advert.Title = advertmodel.Title
		advert.Description = advertmodel.Description
		advert.StartDate = advertmodel.Startdate
		advert.EndDate = advertmodel.Enddate
		advert.MentorID = mentor.ID
		advert.CompanyID = 3

		Config.DB.Create(&advert)
		c.JSON(202, &advert)

	}
	if company.ID != 0 {
		mentor.Major = "to use it"
		c.BindJSON(&advertmodel)
		advert.Title = advertmodel.Title
		advert.Description = advertmodel.Description
		advert.StartDate = advertmodel.Startdate
		advert.EndDate = advertmodel.Enddate
		advert.CompanyID = company.ID
		advert.MentorID = 56

		Config.DB.Create(&advert)
		c.JSON(202, &advert)
	}

}

func GetAdvert(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var advert Models.Advert
	advert = Repositories.GetAdvert(c)
	c.JSON(202, advert)
}

func GetAllAdverts(c *gin.Context) {

	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var advert []Models.Advert
	Config.DB.Find(&advert)
	//advert = Repositories.GetAdverts(c)
	c.JSON(202, advert)
}

func GetAdvertSolo(c *gin.Context) {
	type AdvertModel struct {
		Advert    Models.Advert
		Company   Models.Company
		Mentor    Models.Mentor
		Userprof  Models.UserProfile
		User      Models.User
		Truefalse bool
	}

	var advertjson AdvertModel
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var advert Models.Advert
	Config.DB.Where("id = ?", uint(number)).First(&advert)
	fmt.Println(advert.CompanyID)

	var comp Models.Company
	var ment Models.Mentor
	Config.DB.Where("id = ?", advert.CompanyID).First(&comp)
	Config.DB.Where("id = ?", advert.MentorID).First(&ment)
	advertjson.Truefalse = false

	if comp.ID != 0 {
		advertjson.Truefalse = true
		var user Models.User
		Config.DB.Where("id = ?", comp.UserID).First(&user)
		advertjson.User = user
		var userprof Models.UserProfile
		Config.DB.Where("user_id = ?", comp.UserID).First(&userprof)
		advertjson.Userprof = userprof
	}

	if comp.ID == 3 {
		var user Models.User
		Config.DB.Where("id = ?", ment.UserID).First(&user)
		advertjson.User = user
		var userprof Models.UserProfile
		Config.DB.Where("user_id = ?", ment.UserID).First(&userprof)
		advertjson.Userprof = userprof
	}

	advertjson.Advert = advert
	advertjson.Company = comp
	advertjson.Mentor = ment
	c.JSON(200, advertjson)

}

func GetAdvertsAll(c *gin.Context) {

	type Model struct {
		Advert  Models.Advert
		Company Models.Company
		User    Models.User
		Profile Models.UserProfile
		About   Models.About
		What    bool
	}
	var adverts []Models.Advert

	Config.DB.Where("mentor_id= ?", 56).Find(&adverts)
	var jsonmodel []Model
	var jsonsolo Model
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)
	for _, element := range adverts {

		var comp Models.Company
		Config.DB.Where("id = ?", element.CompanyID).First(&comp)
		var user Models.User
		Config.DB.Where("id = ?", comp.UserID).First(&user)
		var prof Models.UserProfile
		Config.DB.Where("user_id = ?", comp.UserID).First(&prof)
		var about Models.About
		Config.DB.Where("user_id = ?", comp.UserID).First(&about)

		var userrr Models.User
		Config.DB.Where("mail = ?", controll).First(&userrr)
		var app Models.Application
		Config.DB.Where("user_id = ? AND advert_id = ?", userrr.ID, element.ID).First(&app)

		if app.ID != 0 {
			jsonsolo.What = false
		}
		if app.ID == 0 {
			jsonsolo.What = true
		}

		jsonsolo.About = about
		jsonsolo.User = user
		jsonsolo.Profile = prof
		jsonsolo.Advert = element
		jsonsolo.Company = comp
		jsonmodel = append(jsonmodel, jsonsolo)
	}

	c.JSON(200, jsonmodel)
}
