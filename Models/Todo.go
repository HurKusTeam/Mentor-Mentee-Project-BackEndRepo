package Models

type Todo struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	Description   string `gorm:"size:500"`
	Status        int
	EndDate       string
	ActionDate    string
	Issuccessfull bool
	IsDeleted     bool
	MenteeID      uint
	MentorID      uint
	Comments      []Comment `gorm:"foreignKey:TodoID"`
}
