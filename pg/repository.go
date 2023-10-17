package pg

import (
	"gorm.io/gorm"
)

type IRepository interface {
	SaveImage(data []byte) error
	GetImage(id uint) (*ImageModel, error)
	AutoMigrate() error
}

type ImageRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) IRepository {
	return ImageRepository{DB: db}
}

func (r ImageRepository) SaveImage(data []byte) error {
	return r.DB.Create(&ImageModel{Data: data}).Error
}

func (r ImageRepository) GetImage(id uint) (*ImageModel, error) {
	var img *ImageModel
	res := r.DB.Model(ImageModel{}).Where("id = ?", id).First(&img)
	return img, res.Error
}

func (r ImageRepository) AutoMigrate() error {
	return r.DB.AutoMigrate(&ImageModel{})
}
