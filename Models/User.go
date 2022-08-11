package Models

import "time"

type User struct {
	ID           uint `gorm:"primaryKey"`
	UserName     string
	Mail         string
	Password     string
	RegisterDate time.Time
	LastLogin    time.Time
	IsDeleted    bool
	Skills       []Skill       `gorm:"foreignKey:UserID"`
	Languages    []Language    `gorm:"foreignKey:UserID"`
	Applications []Application `gorm:"foreignKey:UserID"`
	Mentors      []Mentor      `gorm:"foreignKey:UserID"`
	Mentees      []Mentee      `gorm:"foreignKey:UserID"`
	Companies    []Company     `gorm:"foreignKey:UserID"`
	UserProfiles []UserProfile `gorm:"foreignKey:UserID"`
	Abouts       []About       `gorm:"foreignKey:UserID"`
	Universities []University  `gorm:"foreignKey:UserID"`
}
