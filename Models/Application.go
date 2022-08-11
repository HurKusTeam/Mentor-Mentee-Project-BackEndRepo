package Models

type Application struct {
	ID        uint `gorm:"primaryKey"`
	Status    int
	IsDeleted bool
	Date      string
	UserID    uint
	//MentorID  uint
	AdvertID uint
}
