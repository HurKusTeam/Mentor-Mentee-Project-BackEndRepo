package Models

type About struct {
	ID        uint   `gorm:"primaryKey"`
	Facebook  string `gorm:"size:1000"`
	Twitter   string `gorm:"size:1000"`
	Linkedin  string `gorm:"size:1000"`
	Website   string `gorm:"size:1000"`
	GitHub    string `gorm:"size:1000"`
	IsDeleted bool
	UserID    uint
}
