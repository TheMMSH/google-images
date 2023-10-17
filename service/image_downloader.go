package service

import (
	"google-images/crypt"
	"google-images/googleapis"
	"google-images/img"
	"google-images/pg"
	"log"
	"math"
	"sync"
)

type IDownloaderService interface {
	ProcessImagesConcurrently(query string, maxImages int)
}

type ImageDownloaderService struct {
	googleApi    googleapis.IGoogleApiService
	imageResizer img.IResizer
	cr           crypt.ICrypt
	repo         pg.IRepository
}

func NewIDownloaderService(googleApi googleapis.IGoogleApiService, imageResizer img.IResizer, encryption crypt.ICrypt, repo pg.IRepository) IDownloaderService {
	return &ImageDownloaderService{
		googleApi:    googleApi,
		imageResizer: imageResizer,
		cr:           encryption,
		repo:         repo,
	}
}

func (d *ImageDownloaderService) ProcessImagesConcurrently(query string, maxImages int) {
	var wg sync.WaitGroup

	downloadedImagesCh := make(chan []byte, maxImages)
	resizedImagesCh := make(chan []byte, maxImages)
	encryptedImagesCh := make(chan []byte, maxImages)

	wg.Add(4)

	// Concurrently Download Images
	go func() {
		defer wg.Done()
		d.downloadImages(query, maxImages, downloadedImagesCh)
	}()

	// Concurrently Resize Images
	go func() {
		defer wg.Done()
		d.resizeImages(downloadedImagesCh, resizedImagesCh)
	}()

	// Concurrently Encrypt Images
	go func() {
		defer wg.Done()
		d.encryptImages(resizedImagesCh, encryptedImagesCh)
	}()

	// Concurrently Store Images to Database
	go func() {
		defer wg.Done()
		d.storeImages(encryptedImagesCh)
	}()

	wg.Wait()
}

func (d *ImageDownloaderService) downloadImages(query string, maxImages int, downloadedImagesCh chan<- []byte) {
	pages := int(math.Ceil(float64(maxImages) / float64(googleapis.GooglePageResultsSize)))
	var dlImagesWG sync.WaitGroup
	for i := 0; i < pages; i++ {
		dlImagesWG.Add(1)
		p := i
		go func() {
			defer dlImagesWG.Done()
			res, err := d.googleApi.DownloadImages(query, p)
			if err != nil {
				log.Println(err)
				return
			}
			for _, re := range res {
				downloadedImagesCh <- re
			}
		}()
	}
	dlImagesWG.Wait()
	close(downloadedImagesCh)
}

func (d *ImageDownloaderService) resizeImages(downloadedImagesCh <-chan []byte, resizedImagesCh chan<- []byte) {
	for image := range downloadedImagesCh {
		resizedImage, err := d.imageResizer.ResizeImage(image)
		if err != nil {
			log.Println(err)
			continue
		}
		resizedImagesCh <- resizedImage
	}
	close(resizedImagesCh)
}

func (d *ImageDownloaderService) encryptImages(resizedImagesCh <-chan []byte, encryptedImagesCh chan<- []byte) {
	for resizedImage := range resizedImagesCh {
		encryptedImage, err := d.cr.Encrypt(resizedImage)
		if err != nil {
			log.Println(err)
			continue
		}
		encryptedImagesCh <- encryptedImage
	}
	close(encryptedImagesCh)
}

func (d *ImageDownloaderService) storeImages(encryptedImagesCh <-chan []byte) {
	for encryptedImage := range encryptedImagesCh {
		err := d.repo.SaveImage(encryptedImage)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
