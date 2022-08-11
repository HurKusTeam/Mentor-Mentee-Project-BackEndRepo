package Models

type LanguageCatalog struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	IsDeleted bool
	Languages []Language `gorm:"foreignKey:LanguageCatalogID"`
}
