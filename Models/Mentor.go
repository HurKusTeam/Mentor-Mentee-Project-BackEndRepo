package Models

type Mentor struct {
	ID           uint `gorm:"primaryKey"`
	Major        string
	IsIndividual bool
	IsDeleted    bool
	UserID       uint
	CompanyID    uint
	//Applications []Application `gorm:"foreignKey:MentorID"`
	Todos    []Todo    `gorm:"foreignKey:MentorID"`
	Comments []Comment `gorm:"foreignKey:MentorID"`
	Meetings []Meeting `gorm:"foreignKey:MentorID"`
	Mentees  []Mentee  `gorm:"foreignKey:MentorID"`
	Adverts  []Advert  `gorm:"foreignKey:MentorID"`
}
