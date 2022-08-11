package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"crypto/sha256"
	"fmt"

	"github.com/gin-gonic/gin"
)

//SESSION DA MAIL TULUNCA HATA VERIYOR  MAIL DEGISMEK ISTENDIGINDE O SEBEPLE SESSION DA ID TUTULMALI!!!

func EditProfile(c *gin.Context) {
	var user Models.User
	var modeluser Models.User
	var userps []Models.UserProfile
	var userp Models.UserProfile
	var mentors []Models.Mentor
	var mentor Models.Mentor
	var mentees []Models.Mentee
	var mentee Models.Mentee
	var companies []Models.Company
	var company Models.Company
	var abouts []Models.About
	var about Models.About

	var universities []Models.University
	var uni Models.University

	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionid"]

	Config.DB.Where("id = ?", control).First(&user)
	c.ShouldBindJSON(&modeluser)

	if len(modeluser.UserName) != 0 {
		user.UserName = modeluser.UserName
	}
	if len(modeluser.Mail) != 0 {
		user.Mail = modeluser.Mail
	}
	if len(modeluser.Password) != 0 {
		user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(modeluser.Password)))
	}

	upc := int(Config.DB.Where("user_id = ?", user.ID).Find(&userps).RowsAffected)
	umentorc := int(Config.DB.Where("user_id = ?", user.ID).Find(&mentors).RowsAffected)
	umenteec := int(Config.DB.Where("user_id = ?", user.ID).Find(&mentees).RowsAffected)
	ucompanyc := int(Config.DB.Where("user_id = ?", user.ID).Find(&companies).RowsAffected)
	uaboutc := int(Config.DB.Where("user_id = ?", user.ID).Find(&abouts).RowsAffected)

	uunic := int(Config.DB.Where("user_id = ?", user.ID).Find(&universities).RowsAffected)

	if int(len(modeluser.UserProfiles)) != 0 {
		if upc == 0 {
			//Config.DB.Model(&user).Association("UserProfiles").Append(&modeluser.UserProfiles)
			user.UserProfiles = append(user.UserProfiles, modeluser.UserProfiles[0])
		}
		if upc == 1 {
			Config.DB.Where("user_id = ?", user.ID).First(&userp)

			userp.Address = modeluser.UserProfiles[0].Address
			userp.Biography = modeluser.UserProfiles[0].Biography
			userp.BirthDate = modeluser.UserProfiles[0].BirthDate
			userp.City = modeluser.UserProfiles[0].City
			userp.Name = modeluser.UserProfiles[0].Name
			userp.PhoneNumber = modeluser.UserProfiles[0].PhoneNumber
			//userp.ProfileImage = modeluser.UserProfiles[0].ProfileImage
			userp.Surname = modeluser.UserProfiles[0].Surname
			Config.DB.Save(&userp)
			c.JSON(200, &userp)
			fmt.Println(modeluser.UserProfiles[0].Name)
			fmt.Println(userp.UserID)
		}
	}

	if int(len(modeluser.Mentors)) != 0 {
		if umentorc == 0 {
			user.Mentors = append(user.Mentors, modeluser.Mentors[0])
		}
		if umentorc == 1 {
			Config.DB.Where("user_id = ?", user.ID).First(&mentor)
			mentor.Major = modeluser.Mentors[0].Major
			Config.DB.Save(&mentor)
			c.JSON(200, &mentor)
		}
	}

	if int(len(modeluser.Mentees)) != 0 {
		if umenteec == 0 {
			var m Models.Mentee
			m.Department = modeluser.Mentees[0].Department
			m.GPA = modeluser.Mentees[0].GPA
			Config.DB.Create(&m)
			user.Mentees = append(user.Mentees, modeluser.Mentees[0])
		}
		if umenteec == 1 {
			Config.DB.Where("user_id = ?", user.ID).First(&mentee)
			mentee.Department = modeluser.Mentees[0].Department
			mentee.GPA = modeluser.Mentees[0].GPA
			Config.DB.Save(&mentee)
			c.JSON(200, &mentee)
		}
	}
	if int(len(modeluser.Companies)) != 0 {
		if ucompanyc == 0 {
			user.Companies = append(user.Companies, modeluser.Companies[0])
		}
		if ucompanyc == 1 {
			Config.DB.Where("user_id = ?", user.ID).First(&company)
			company.Description = modeluser.Companies[0].Description
			company.PersonalCount = modeluser.Companies[0].PersonalCount
			company.Sector = modeluser.Companies[0].Sector
			company.SinceDate = modeluser.Companies[0].SinceDate
			company.Title = modeluser.Companies[0].Title
			company.Type = modeluser.Companies[0].Type

			Config.DB.Save(&company)
			c.JSON(200, &company)
		}
	}

	if int(len(modeluser.Abouts)) != 0 {
		if uaboutc == 0 {
			user.Abouts = append(user.Abouts, modeluser.Abouts[0])
		}
		if uaboutc == 1 {
			Config.DB.Where("user_id = ?", user.ID).First(&about)
			about.Facebook = modeluser.Abouts[0].Facebook
			about.GitHub = modeluser.Abouts[0].GitHub
			about.Linkedin = modeluser.Abouts[0].Linkedin
			about.Twitter = modeluser.Abouts[0].Twitter
			about.Website = modeluser.Abouts[0].Website
			Config.DB.Save(&about)
			c.JSON(200, &about)
		}
	}
	if int(len(modeluser.Universities)) != 0 {
		if uunic == 0 {
			user.Universities = append(user.Universities, modeluser.Universities[0])
		}
		if uunic == 1 {
			Config.DB.Where("user_id = ?", user.ID).First(&uni)
			uni.UniversityCatalogID = modeluser.Universities[0].UniversityCatalogID
			Config.DB.Save(&uni)
			c.JSON(200, &uni)

		}
	}

	c.ShouldBindJSON(&user)
	Config.DB.Save(&user)
	c.JSON(200, &user)
}

type ModelE struct {
	UserName       string
	Mail           string
	Password       string
	Name           string
	Surname        string
	Biography      string
	BirthDate      string
	PhoneNumber    string
	ProfileImage   string
	City           string
	Address        string
	University     string
	Major          string
	Department     string
	IsIndividual   bool
	Badge          uint
	GPA            float64
	MenteeCount    int
	Skills         []string
	Title          string
	CompanyID      uint
	Description    string
	Sector         string
	PersonalCount  int
	SinceDate      string
	Type           string
	Facebook       string
	Twitter        string
	Linkedin       string
	Website        string
	GitHub         string
	AdvertID       uint
	IsMentor       bool
	IsMentee       bool
	IsCompany      bool
	Languages      []string
	Companymentors []Models.Mentor
	Companyadverts []Models.Advert
}

func GetProfile(c *gin.Context) {

	var userr Models.User
	var userp Models.UserProfile

	var uabout Models.About

	var uapp Models.Advert
	var uuniversity Models.University
	var unic Models.UniversityCatalog
	var langs []Models.Language
	var skills []Models.Skill
	var umentee Models.Mentee
	var umentor Models.Mentor
	var ucompany Models.Company

	var comad []Models.Advert
	var skillname []string
	var languagename []string
	var mentorscompany Models.Company
	var commen []Models.Mentor
	var model ModelE

	session, _ := store.Get(c.Request, "sessioncontrol")
	i := session.Values["sessionid"]
	Config.DB.First(&userr, "id = ?", i)
	Config.DB.First(&userp, "user_id = ?", i)
	skillcount := int(Config.DB.Find(&skills, "user_id = ?", i).RowsAffected)
	Config.DB.First(&uabout, "user_id = ?", i)
	langcount := int(Config.DB.Find(&langs, "user_id = ?", i).RowsAffected)

	Config.DB.First(&uuniversity, "user_id = ?", i)
	Config.DB.First(&unic, "id = ?", uuniversity.UniversityCatalogID)
	Config.DB.First(&umentee, "user_id = ?", i)
	Config.DB.First(&umentor, "user_id = ?", i)
	Config.DB.First(&ucompany, "user_id = ?", i)
	Config.DB.First(&uapp, "mentor_id = ?", umentor.ID)

	model.IsCompany = false
	model.IsMentee = false
	model.IsMentor = false
	if umentee.ID != 0 {
		model.IsMentee = true
	}
	if ucompany.ID != 0 {
		model.IsCompany = true
		model.Title = ucompany.Title
		model.Sector = ucompany.Sector
		model.Description = ucompany.Description
		model.SinceDate = ucompany.SinceDate
		model.PersonalCount = ucompany.PersonalCount
		model.Type = ucompany.Type
		Config.DB.Where("company_id = ?", ucompany.ID).Find(&commen)
		model.Companymentors = commen
		Config.DB.Where("company_id = ?", ucompany.ID).Find(&comad)
		model.Companyadverts = comad

	}
	if umentor.ID != 0 {
		model.IsMentor = true
		Config.DB.First(&mentorscompany, "id = ?", umentor.CompanyID)
		model.Title = mentorscompany.Title
		model.CompanyID = mentorscompany.ID
		model.Sector = mentorscompany.Sector
		model.Description = mentorscompany.Description
		model.SinceDate = mentorscompany.SinceDate
		model.PersonalCount = mentorscompany.PersonalCount
		model.Type = mentorscompany.Type
		if umentor.CompanyID == 3 {
			model.IsIndividual = true
		}
	}

	model.UserName = userr.UserName
	model.Mail = userr.Mail
	fmt.Printf(userr.Mail)
	model.Password = userr.Password
	model.Name = userp.Name
	model.Surname = userp.Surname
	model.Address = userp.Address
	model.City = userp.City
	model.Biography = userp.Biography
	model.BirthDate = userp.BirthDate
	model.PhoneNumber = userp.PhoneNumber
	model.ProfileImage = userp.ProfileImage

	model.Major = umentor.Major
	model.MenteeCount = umentee.MenteeCount
	model.GPA = umentee.GPA
	model.Department = umentee.Department
	model.Badge = umentee.Badge

	model.AdvertID = uapp.ID
	model.Facebook = uabout.Facebook
	model.Twitter = uabout.Twitter
	model.GitHub = uabout.GitHub
	model.Linkedin = uabout.Linkedin
	model.Website = uabout.Website
	model.University = unic.Name
	fmt.Println(langcount)
	fmt.Println(skillcount)
	for _, lang := range langs {
		var langc Models.LanguageCatalog
		var ln string
		fmt.Printf("lang:")
		fmt.Println(langcount)
		Config.DB.First(&langc, "id = ?", lang.LanguageCatalogID)
		fmt.Println(ln)
		fmt.Printf("catalogname:")
		fmt.Println(langc.Name)
		ln = langc.Name
		languagename = append(languagename, langc.Name)
	}

	for _, skill := range skills {
		var skillc Models.SkillCatalog
		var sn string
		fmt.Printf("skill:")
		fmt.Println(skillcount)
		Config.DB.First(&skillc, "id = ?", skill.SkillCatalogID)
		fmt.Println(sn)
		fmt.Printf("catalognames:")
		fmt.Println(skillc.Name)
		sn = skillc.Name
		skillname = append(skillname, skillc.Name)
	}

	for _, snm := range skillname {
		fmt.Println(snm)
	}
	for _, lnm := range languagename {
		fmt.Println(lnm)
	}
	model.Skills = skillname
	model.Languages = languagename

	for _, skname := range model.Skills {
		fmt.Println(skname)
	}
	for _, lname := range model.Languages {
		fmt.Println(lname)
	}

	err := c.ShouldBindJSON(&model)
	c.JSON(200, model)

	fmt.Print(err)
	return
}

func DeleteSkill(c *gin.Context) {
	var skillc Models.SkillCatalog
	var skill Models.Skill

	session, _ := store.Get(c.Request, "sessioncontrol")
	i := session.Values["sessionid"]

	Config.DB.Where("name = ?", c.Param("name")).First(&skillc)
	Config.DB.Where("user_id = ? AND skill_catalog_id = ?", i, skillc.ID).Delete(&skill)
}
func DeleteLanguage(c *gin.Context) {
	var langc Models.LanguageCatalog
	var lang Models.Language

	session, _ := store.Get(c.Request, "sessioncontrol")
	i := session.Values["sessionid"]

	Config.DB.Where("name = ?", c.Param("name")).First(&langc)
	Config.DB.Where("user_id = ? AND language_catalog_id = ?", i, langc.ID).Delete(&lang)
}

type UserImage struct {
	Image string
}

func ImageUpdate(c *gin.Context) {

	var image UserImage
	var userprofile Models.UserProfile

	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionid"]

	if control == nil {
		fmt.Println("SESSİON BOŞ")
		c.Redirect(301, "/Login")
	}

	c.BindJSON(&image)
	Config.DB.First(&userprofile, "user_id = ?", control)

	if userprofile.UserID != control {
		c.JSON(400, "kullanıcı bulunamadı.")

	} else {
		userprofile.ProfileImage = image.Image
		Config.DB.Save(&userprofile)
	}

}

type SkillIDArray struct {
	Skillids []uint
}

func UpdateSkill(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	i := session.Values["sessionid"]
	var skills []Models.Skill
	var sids SkillIDArray
	var user Models.User

	var skillfound bool
	c.ShouldBindJSON(&sids)
	Config.DB.Where("user_id = ?", i).Find(&skills)
	Config.DB.First(&user, "id = ?", i)
	if len(sids.Skillids) != 0 {
		for _, mskill := range sids.Skillids {
			var skill Models.Skill
			skillfound = false
			for _, skv := range skills {
				if skv.SkillCatalogID == mskill {
					skillfound = true
				}
			}
			if !skillfound {
				skill.UserID = user.ID
				skill.SkillCatalogID = mskill
				Config.DB.Create(&skill)
				user.Skills = append(user.Skills, skill)
			}
		}

	}

}

type LanguageIDArray struct {
	Langids []uint
}

func UpdateLang(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	i := session.Values["sessionid"]
	var languages []Models.Language
	var lids LanguageIDArray
	var langfound bool
	var user Models.User
	c.ShouldBindJSON(&lids)
	Config.DB.Where("user_id = ?", i).Find(&languages)
	Config.DB.First(&user, "id = ?", i)
	if len(lids.Langids) != 0 {

		for _, mlang := range lids.Langids {
			langfound = false
			fmt.Println(mlang)
			var lang Models.Language
			for _, ulan := range languages {
				if mlang == ulan.LanguageCatalogID {
					langfound = true
					fmt.Println("bulundu")
				}

			}
			if !langfound {
				lang.LanguageCatalogID = mlang
				lang.UserID = user.ID
				Config.DB.Create(&lang)
				user.Languages = append(user.Languages, lang)
			}

		}
	}

}

func GetProfileElse(c *gin.Context) {
	var userr Models.User
	var userp Models.UserProfile

	var uabout Models.About

	var uapp Models.Application
	var uuniversity Models.University
	var unic Models.UniversityCatalog
	var langs []Models.Language
	var skills []Models.Skill
	var umentee Models.Mentee
	var umentor Models.Mentor
	var ucompany Models.Company

	var comad []Models.Advert
	var skillname []string
	var languagename []string
	var mentorscompany Models.Company
	var commen []Models.Mentor
	var model ModelE

	i := c.Param("id")
	Config.DB.First(&userr, "id = ?", i)
	Config.DB.First(&userp, "user_id = ?", i)
	skillcount := int(Config.DB.Find(&skills, "user_id = ?", i).RowsAffected)
	Config.DB.First(&uabout, "user_id = ?", i)
	langcount := int(Config.DB.Find(&langs, "user_id = ?", i).RowsAffected)
	Config.DB.First(&uapp, "user_id = ?", i)
	Config.DB.First(&uuniversity, "user_id = ?", i)
	Config.DB.First(&unic, "id = ?", uuniversity.UniversityCatalogID)
	Config.DB.First(&umentee, "user_id = ?", i)
	Config.DB.First(&umentor, "user_id = ?", i)
	Config.DB.First(&ucompany, "user_id = ?", i)

	model.IsCompany = false
	model.IsMentee = false
	model.IsMentor = false
	if umentee.ID != 0 {
		model.IsMentee = true
	}
	if ucompany.ID != 0 {
		model.IsCompany = true
		model.Title = ucompany.Title
		model.Sector = ucompany.Sector
		model.Description = ucompany.Description
		model.SinceDate = ucompany.SinceDate
		model.PersonalCount = ucompany.PersonalCount
		model.Type = ucompany.Type
		Config.DB.Where("company_id = ?", ucompany.ID).Find(&commen)
		model.Companymentors = commen
		Config.DB.Where("company_id = ?", ucompany.ID).Find(&comad)
		model.Companyadverts = comad

	}
	if umentor.ID != 0 {
		model.IsMentor = true
		Config.DB.First(&mentorscompany, "id = ?", umentor.CompanyID)
		model.Title = mentorscompany.Title
		model.CompanyID = mentorscompany.ID
		model.Sector = mentorscompany.Sector
		model.Description = mentorscompany.Description
		model.SinceDate = mentorscompany.SinceDate
		model.PersonalCount = mentorscompany.PersonalCount
		model.Type = mentorscompany.Type
	}

	model.UserName = userr.UserName
	model.Mail = userr.Mail
	model.Password = userr.Password
	model.Name = userp.Name
	model.Surname = userp.Surname
	model.Address = userp.Address
	model.City = userp.City
	model.Biography = userp.Biography
	model.BirthDate = userp.BirthDate
	model.PhoneNumber = userp.PhoneNumber
	model.ProfileImage = userp.ProfileImage

	model.Major = umentor.Major
	model.MenteeCount = umentee.MenteeCount
	model.GPA = umentee.GPA
	model.Department = umentee.Department
	model.Badge = umentee.Badge

	model.AdvertID = uapp.AdvertID
	model.Facebook = uabout.Facebook
	model.Twitter = uabout.Twitter
	model.GitHub = uabout.GitHub
	model.Linkedin = uabout.Linkedin
	model.Website = uabout.Website
	model.University = unic.Name
	fmt.Println(langcount)
	fmt.Println(skillcount)
	for _, lang := range langs {
		var langc Models.LanguageCatalog
		var ln string
		fmt.Printf("lang:")
		fmt.Println(langcount)
		Config.DB.First(&langc, "id = ?", lang.LanguageCatalogID)
		fmt.Println(ln)
		fmt.Printf("catalogname:")
		fmt.Println(langc.Name)
		ln = langc.Name
		languagename = append(languagename, langc.Name)
	}

	for _, skill := range skills {
		var skillc Models.SkillCatalog
		var sn string
		fmt.Printf("skill:")
		fmt.Println(skillcount)
		Config.DB.First(&skillc, "id = ?", skill.SkillCatalogID)
		fmt.Println(sn)
		fmt.Printf("catalognames:")
		fmt.Println(skillc.Name)
		sn = skillc.Name
		skillname = append(skillname, skillc.Name)
	}

	for _, snm := range skillname {
		fmt.Println(snm)
	}
	for _, lnm := range languagename {
		fmt.Println(lnm)
	}
	model.Skills = skillname
	model.Languages = languagename

	for _, skname := range model.Skills {
		fmt.Println(skname)
	}
	for _, lname := range model.Languages {
		fmt.Println(lname)
	}

	err := c.ShouldBindJSON(&model)
	c.JSON(200, model)

	fmt.Print(err)
	return
}

type SkillAdded struct {
	Name string
}

func AddSkill(c *gin.Context) {
	var sadd SkillAdded
	var sk Models.SkillCatalog
	var skills []Models.SkillCatalog
	var found bool
	found = false
	c.BindJSON(&sadd)
	Config.DB.Find(&skills)
	for _, sa := range skills {
		if sa.Name == sadd.Name {
			found = true
		}
	}
	if !found {
		sk.Name = sadd.Name
		Config.DB.Create(&sk)
		c.JSON(200, sk)
	}
}
