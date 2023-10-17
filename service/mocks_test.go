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

func (m *MockIRepository) AutoMigrate() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockIRepository) SaveImage(data []byte) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockIRepository) GetImage(id uint) (*pg.ImageModel, error) {
	args := m.Called(id)
	return args.Get(0).(*pg.ImageModel), args.Error(1)
}

type MockICrypt struct {
	mock.Mock
}

func (m *MockICrypt) Encrypt(plaintext []byte) ([]byte, error) {
	args := m.Called(plaintext)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockICrypt) Decrypt(ciphertext []byte) ([]byte, error) {
	args := m.Called(ciphertext)
	return args.Get(0).([]byte), args.Error(1)
}

type MockIResizer struct {
	mock.Mock
}

func (m *MockIResizer) ResizeImage(data []byte) ([]byte, error) {
	args := m.Called(data)
	return args.Get(0).([]byte), args.Error(1)
}
