package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"TREgitim/Repositories"
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisteringUser struct {
	Username string
	Mail     string
	Password string
	Dropdown int
}

func Register(c *gin.Context) {
	var user Models.User
	var reguser RegisteringUser
	var userr Models.User
	var company Models.Company
	//var mentee Models.Mentee
	var mentor Models.Mentor

	c.BindJSON(&reguser)
	Config.DB.First(&user, "mail=?", reguser.Mail)
	//Config.DB.First(&mentor, "id=?", user.ID)

	if user.Mail == reguser.Mail {
		c.JSON(400, "Bu mail adresi kullanılıyor.")
		c.Redirect(http.StatusFound, "/Login")
		//c.Redirect(301, "/Login")
	}
	if user.Mail != reguser.Mail {
		password := fmt.Sprintf("%x", sha256.Sum256([]byte(reguser.Password)))
		user.Password = password
		user.UserName = reguser.Username
		user.Mail = reguser.Mail
		user.RegisterDate = time.Now()
		Repositories.NewUser(user)

		Config.DB.Where("mail = ?", user.Mail).First(&userr)
		if reguser.Dropdown == 1 {
			mentor.UserID = userr.ID
			//mentor.IsIndividual = true
			mentor.CompanyID = 3
			Config.DB.Create(&mentor)
			c.JSON(202, mentor)
			c.Redirect(http.StatusFound, "/MentorPage")

		}
		if reguser.Dropdown == 2 {
			company.UserID = userr.ID
			Config.DB.Create(&company)
			c.JSON(202, "şirket olarak kaydınız yapıldı")
			c.Redirect(http.StatusFound, "/CompanyPage")

		}
		c.JSON(202, "kaydınız yapıldı")

	}

}
