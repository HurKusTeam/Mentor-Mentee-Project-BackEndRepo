package Models

type Company struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	Description   string `gorm:"size:1000"`
	Sector        string
	PersonalCount int
	SinceDate     string
	Type          string
	IsDeleted     bool
	UserID        uint
	Mentees       []Mentee `gorm:"foreignKey:CompanyID"`
	Mentors       []Mentor `gorm:"foreignKey:CompanyID"`
	Adverts       []Advert `gorm:"foreignKey:CompanyID"`
}
