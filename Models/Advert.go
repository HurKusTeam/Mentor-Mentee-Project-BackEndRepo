package Models

type Advert struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Description  string `gorm:"size:1000"`
	StartDate    string
	EndDate      string
	IsDeleted    bool
	CompanyID    uint
	MentorID     uint
	Applications []Application `gorm:"foreignKey:AdvertID"`
}
