package service

import (
	"github.com/stretchr/testify/mock"
	"google-images/googleapis"
	"testing"
)

func TestProcessImagesConcurrentlyCallsDownloadAndResizeAndEncryptAndSave(t *testing.T) {
	mockCrypt := new(MockICrypt)
	mockResizer := new(MockIResizer)
	mockRepository := new(MockIRepository)
	mockGoogleApiService := new(MockIGoogleApiService)

	sut := NewIDownloaderService(mockGoogleApiService, mockResizer, mockCrypt, mockRepository)

	mockCrypt.On("Encrypt", mock.Anything).Return([]byte("cipheredText"), nil)
	mockResizer.On("ResizeImage", mock.Anything).Return([]byte("resizedImage"), nil)
	mockRepository.On("SaveImage", mock.Anything).Return(nil)
	mockGoogleApiService.On("DownloadImages", mock.Anything, mock.Anything).Return([]googleapis.MemImage{[]byte("img1")}, nil)

	sut.ProcessImagesConcurrently("testQuery", 100)

	mockGoogleApiService.AssertCalled(t, "DownloadImages", "testQuery", 9)
	mockResizer.AssertCalled(t, "ResizeImage", []byte("img1"))
	mockCrypt.AssertCalled(t, "Encrypt", []byte("resizedImage"))
	mockRepository.AssertCalled(t, "SaveImage", []byte("cipheredText"))
}
