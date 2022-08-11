package Models

type SkillCatalog struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Skills []Skill `gorm:"foreignKey:SkillCatalogID"`
}
