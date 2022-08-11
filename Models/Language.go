package Models

type Language struct {
	ID                uint `gorm:"primaryKey"`
	IsDeleted         bool
	LanguageCatalogID uint
	UserID            uint
}
