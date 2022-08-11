package Models

type UniversityCatalog struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Universities []University `gorm:"foreignKey:UniversityCatalogID"`
}
