package service

import (
	"google-images/crypt"
	"google-images/pg"
)

type IPresenterService interface {
	ViewImage(id uint) ([]byte, error)
}

type PresenterService struct {
	cr   crypt.ICrypt
	repo pg.IRepository
}

func NewIPresenterService(encryption crypt.ICrypt, repo pg.IRepository) IPresenterService {
	return PresenterService{
		cr:   encryption,
		repo: repo,
	}
}

func (p PresenterService) ViewImage(id uint) ([]byte, error) {
	img, err := p.repo.GetImage(id)
	if err != nil {
		return nil, err
	}

	return p.cr.Decrypt(img.Data)
}
