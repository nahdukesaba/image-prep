package manager

import (
	"image-prep/service"
	"sync"
)

var (
	imageServiceOnce sync.Once
	imageService     service.ImageService
)

func ImageService() service.ImageService {
	imageServiceOnce.Do(func() {
		imageService = service.NewImageService()
	})

	return imageService
}
