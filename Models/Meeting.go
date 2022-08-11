package Models

type Meeting struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string `gorm:"size:500"`
	Link        string `gorm:"size:500"`
	Date        string
	IsDeleted   bool
	MenteeID    uint
	MentorID    uint
}
