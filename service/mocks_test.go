package service

import (
	"github.com/stretchr/testify/mock"
	"google-images/googleapis"
	"google-images/pg"
)

type MockIGoogleApiService struct {
	mock.Mock
}

func (m *MockIGoogleApiService) DownloadImages(query string, page int) ([]googleapis.MemImage, error) {
	args := m.Called(query, page)
	return args.Get(0).([]googleapis.MemImage), args.Error(1)
}

type MockIRepository struct {
	mock.Mock
}

func (m *MockIRepository) StoreImage(data []byte) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockIRepository) FetchImage(id uint) (*pg.ImageModel, error) {
	args := m.Called(id)
	return args.Get(0).(*pg.ImageModel), args.Error(1)
}
