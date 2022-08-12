package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type MailPW struct {
	Mail        string
	Password    string
	Role        int
	UserID      uint
	CompanyID   uint
	MentorID    uint
	MenteeID    uint
	AdvertID    uint
	ProfilIMG   string
	MenteeCount int
	IsInd       bool
}

var store = sessions.NewCookieStore([]byte("sessioncontrol"))

func Login(c *gin.Context) {

	var user MailPW
	var userdb Models.User
	var mentor Models.Mentor
	var mentee Models.Mentee
	var company Models.Company
	var advert Models.Advert
	var userp Models.UserProfile

	c.BindJSON(&user)
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))
	Config.DB.First(&userdb, "mail = ?", user.Mail)
	Config.DB.First(&userp, "user_id = ?", userdb.ID)
	var sesm = user.Mail
	var sid = userdb.ID

	if userdb.Mail != user.Mail {
		c.JSON(400, "Sistemde kullanıcı bulunamadı.")
		c.Redirect(301, "/Login")
	}
	if userdb.Mail == user.Mail && userdb.Password != password {
		c.JSON(400, "Hatalı şifre girdiniz.")
		c.Redirect(301, "/Login")
	}
	if userdb.Mail == user.Mail && userdb.Password == password {

		session, _ := store.Get(c.Request, "sessioncontrol")
		session.Values["sessionmail"] = sesm
		session.Values["sessionid"] = sid
		session.Save(c.Request, c.Writer)
		user.UserID = userdb.ID

		var userr Models.User
		var childmentee Models.Childmentee
		Config.DB.Where("mail = ?", sesm).First(&userr)
		Config.DB.Where("user_id = ?", userr.ID).First(&childmentee)
		if childmentee.Active != true && childmentee.What != true {
			var mentee Models.Mentee
			Config.DB.Where("id = ?", childmentee.MenteeID).First(&mentee)
			mentee.MenteeCount += 1
			childmentee.Active = true
			childmentee.What = true
			Config.DB.Save(&mentee)
			Config.DB.Save(&childmentee)
		}

		Config.DB.First(&mentor, "user_id", userdb.ID)
		Config.DB.First(&mentee, "user_id", userdb.ID)
		Config.DB.First(&company, "user_id", userdb.ID)
		if mentor.ID != 0 {
			user.Role = 0
			user.CompanyID = mentor.CompanyID
			user.MentorID = mentor.ID
			Config.DB.First(&advert, "mentor_id", mentor.ID)
			if advert.ID == 0 {
				user.AdvertID = 0
			}
			if advert.ID != 0 {
				user.AdvertID = advert.ID
			}
			if mentor.IsIndividual {
				user.IsInd = true
			}
		}
		if mentee.ID != 0 {
			user.Role = 1
			user.CompanyID = mentee.CompanyID
			user.MentorID = mentee.MentorID
			user.MenteeID = mentee.ID
			user.MenteeCount = mentee.MenteeCount
		}
		if company.ID != 0 {
			user.Role = 2
			user.CompanyID = company.ID

		}
		if company.ID == 0 && mentee.ID == 0 && mentor.ID == 0 {
			user.Role = 3
		}
		if userp.ID != 0 {
			user.ProfilIMG = userp.ProfileImage
		}
		c.JSON(202, user)

	}

}

func Logout(c *gin.Context) {
	session, err := store.Get(c.Request, "sessioncontrol")
	if err != nil {
		fmt.Println(err)
	}
	session.Values["sessionmail"] = ""
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	//deneme := session.Values["sessionmail"]
	//fmt.Println(deneme)
	c.Redirect(200, "/Login")
}

func Logoutt(c *gin.Context) {
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var user Models.User
	Config.DB.Where("id = ?", uint(number)).First(&user)
	user.IsDeleted = false
	Config.DB.Save(&user)
	c.JSON(200, "Çıkış başarılı.")
}
