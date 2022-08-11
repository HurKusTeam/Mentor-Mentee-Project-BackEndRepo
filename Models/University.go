package Models

type University struct {
	ID                  uint `gorm:"primaryKey"`
	IsDeleted           bool
	UserID              uint
	UniversityCatalogID uint
}
