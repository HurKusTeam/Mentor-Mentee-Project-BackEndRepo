package Models

type Skill struct {
	ID             uint `gorm:"primaryKey"`
	IsDeleted      bool
	SkillCatalogID uint
	UserID         uint
}
