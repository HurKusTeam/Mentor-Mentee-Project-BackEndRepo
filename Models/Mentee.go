package Models

type Mentee struct {
	ID          uint `gorm:"primaryKey"`
	Department  string
	Badge       uint
	GPA         float64 // hata olabilir
	MenteeCount int
	IsDeleted   bool
	UserID      uint
	MentorID    uint
	CompanyID   uint
	Todos       []Todo    `gorm:"foreignKey:MenteeID"`
	Comments    []Comment `gorm:"foreignKey:MenteeID"`
	Meetings    []Meeting `gorm:"foreignKey:MenteeID"`
}
