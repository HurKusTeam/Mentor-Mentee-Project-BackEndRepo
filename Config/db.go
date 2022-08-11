package Config

import (
	"TREgitim/Models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=93.115.79.25 user=elif password=hurKUS.18072022 dbname=hurkus port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("DATABASE CONNECTION HATA")
	}
	db.AutoMigrate(&Models.Childmentee{})
	//db.AutoMigrate(&Models.User{})
	//db.AutoMigrate(&Models.About{})
	//db.AutoMigrate(&Models.Advert{})
	//db.AutoMigrate(&Models.LanguageCatalog{})
	//db.AutoMigrate(&Models.SkillCatalog{})
	//db.AutoMigrate(&Models.UniversityCatalog{})
	//db.AutoMigrate(&Models.Application{})
	//db.AutoMigrate(&Models.Comment{})
	//db.AutoMigrate(&Models.Company{})
	//db.AutoMigrate(&Models.Language{})
	//db.AutoMigrate(&Models.Meeting{})
	//db.AutoMigrate(&Models.Mentee{})
	//db.AutoMigrate(&Models.Mentor{})
	//db.AutoMigrate(&Models.Skill{})
	//db.AutoMigrate(&Models.Todo{})
	//db.AutoMigrate(&Models.University{})
	//db.AutoMigrate(&Models.UserProfile{})
	//db.AutoMigrate(&Models.User{})

	//db.Migrator().DropColumn(&Models.Childmentee{}, "Didcount")
	//db.Migrator().DropColumn(&Models.Application{}, "MentorID")
	//db.Migrator().DropConstraint(&Models.Mentor{}, "Applications")
	//db.Migrator().DropConstraint(&Models.Mentor{}, "fk_mentors_applications")
	//db.Migrator().DropColumn(&Models.Skill{}, "Active")

	//db.Migrator().DropTable(&Models.User{})
	//db.Migrator().DropTable(&Models.About{})
	//db.Migrator().DropTable(&Models.Advert{})
	//db.Migrator().DropTable(&Models.Application{})
	//db.Migrator().DropTable(&Models.Comment{})
	//db.Migrator().DropTable(&Models.Company{})
	//db.Migrator().DropTable(&Models.Language{})
	//db.Migrator().DropTable(&Models.Meeting{})
	//db.Migrator().DropTable(&Models.Mentee{})
	//db.Migrator().DropTable(&Models.Mentor{})
	//db.Migrator().DropTable(&Models.Skill{})
	//db.Migrator().DropTable(&Models.SkillTotal{})
	//db.Migrator().DropTable(&Models.Todo{})
	//db.Migrator().DropTable(&Models.University{})
	//db.Migrator().DropTable(&Models.UniversityTotal{})
	//db.Migrator().DropTable(&Models.UserProfile{})
	//db.Migrator().DropTable(&Models.Childmentee{})

	DB = db
}
