package Models

type UserProfile struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Surname      string
	Biography    string `gorm:"size:500"`
	BirthDate    string
	PhoneNumber  string
	ProfileImage string `gorm:"size:1000"`
	City         string
	Address      string `gorm:"size:500"`
	IsDeleted    bool
	UserID       uint
}
