package Models

type Comment struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"size:500"`
	IsDeleted   bool
	MenteeID    uint
	MentorID    uint
	TodoID      uint
}
