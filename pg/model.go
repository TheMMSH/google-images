package pg

type ImageModel struct {
	ID   uint `gorm:"primaryKey"`
	Data []byte
}

func (m *ImageModel) TableName() string {
	return "images"
}
