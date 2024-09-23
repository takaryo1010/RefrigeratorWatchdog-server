package repository

import (
	"RefrigeratorWatchdog-server/model"
	"io"
	"os"
)

type IImageRepository interface {
	UploadImage(image *model.Image) (*model.Image, error)
}

type imageRepository struct {
}

func NewImageRepository() IImageRepository {
	return &imageRepository{}
}

func (ir *imageRepository) UploadImage(image *model.Image) (*model.Image, error) {

	dst, err := os.Create("images/" + image.Filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, image.ImageFile); err != nil {
		return image, err
	}

	return image, nil
}
