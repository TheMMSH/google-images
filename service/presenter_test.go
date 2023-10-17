package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google-images/pg"
	"testing"
)

func TestPresenterService_ViewImageGetsImageFromDBAndDecryptsImage(t *testing.T) {
	asserts := assert.New(t)
	mockCrypt := new(MockICrypt)
	mockRepository := new(MockIRepository)

	sut := NewIPresenterService(mockCrypt, mockRepository)

	mockCrypt.On("Decrypt", mock.Anything).Return([]byte("plainImage"), nil)
	mockRepository.On("GetImage", mock.Anything).Return(&pg.ImageModel{ID: 10, Data: []byte("cipheredImage")}, nil)

	res, _ := sut.ViewImage(10)

	mockCrypt.AssertCalled(t, "Decrypt", []byte("cipheredImage"))
	mockRepository.AssertCalled(t, "GetImage", uint(10))

	asserts.Equal([]byte("plainImage"), res)
}
