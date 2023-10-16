package pg

import (
	"gorm.io/gorm"
)

type IRepository interface {
	StoreImage(data []byte) error
	FetchImage(id uint) (*ImageModel, error)
}

type ImageRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) IRepository {
	return ImageRepository{DB: db}
}

func (r ImageRepository) StoreImage(data []byte) error {
	return r.DB.Create(&ImageModel{Data: data}).Error
}

func (r ImageRepository) FetchImage(id uint) (*ImageModel, error) {
	var img *ImageModel
	res := r.DB.First(img, id)
	return img, res.Error
}
