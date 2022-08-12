package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"

	"github.com/gin-gonic/gin"
)

type model struct {
	Name       string
	Surname    string
	Major      string
	Uni        string
	Birthdate  string
	Mail       string
	Tel        string
	City       string
	Linkedin   string
	Github     string
	Website    string
	Facebook   string
	Twitter    string
	ID         uint
	AdID       uint
	ProfileImg string
	Applied    bool
}

func GetIndividual(c *gin.Context) {
	var indimentors []Models.Mentor
	var modelmentor model
	var model []model

	Config.DB.Where("is_individual = ?", true).Find(&indimentors)
	for _, ment := range indimentors {
		var up Models.UserProfile
		var ab Models.About
		var us Models.User
		var un Models.University
		var unc Models.UniversityCatalog
		var mentor Models.Mentor
		var advert Models.Advert
		var app Models.Application
		session, _ := store.Get(c.Request, "sessioncontrol")
		control := session.Values["sessionid"]

		Config.DB.Where("user_id = ?", ment.UserID).First(&up)
		Config.DB.Where("user_id = ?", ment.UserID).First(&ab)
		Config.DB.Where("id = ?", ment.UserID).First(&us)
		Config.DB.Where("user_id = ?", ment.UserID).First(&un)
		Config.DB.Where("id = ?", un.UniversityCatalogID).First(&unc)
		Config.DB.Where("user_id = ?", ment.UserID).First(&mentor)
		Config.DB.Where("mentor_id = ?", mentor.ID).First(&advert)
		modelmentor.Name = up.Name
		modelmentor.Surname = up.Surname
		modelmentor.City = up.City
		modelmentor.Birthdate = up.BirthDate
		modelmentor.Tel = up.PhoneNumber
		modelmentor.Mail = us.Mail
		modelmentor.Uni = unc.Name
		modelmentor.Major = ment.Major
		modelmentor.Linkedin = ab.Linkedin
		modelmentor.Facebook = ab.Facebook
		modelmentor.Github = ab.GitHub
		modelmentor.Website = ab.Website
		modelmentor.Twitter = ab.Twitter
		modelmentor.ID = mentor.UserID
		modelmentor.AdID = advert.ID

		modelmentor.Applied = false
		Config.DB.Where("advert_id = ? AND user_id = ?", advert.ID, control).First(&app)
		if app.ID == 0 {
			modelmentor.Applied = false
		}
		if app.ID != 0 {
			modelmentor.Applied = true
		}
		modelmentor.ProfileImg = up.ProfileImage
		model = append(model, modelmentor)

	}
	c.JSON(200, model)

}
