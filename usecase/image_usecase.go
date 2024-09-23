package usecase

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/repository"
	"errors"
	"strconv"
	"time"
)

type IImageUsecase interface {
	UploadImage(file model.Image) (*model.Image, error)
	FetchImage(imageURL string) (*model.Image, error)
}

type imageUsecase struct {
	ir repository.IImageRepository
}

func NewImageUsecase(ir repository.IImageRepository) IImageUsecase {
	return &imageUsecase{ir}
}

func (iu *imageUsecase) UploadImage(file model.Image) (*model.Image, error) {
	if file.ImageFile == nil {
		return nil, errors.New("no image file")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	file.Filename = timestamp + "_" + file.Filename

	return iu.ir.UploadImage(&file)
}

func (iu *imageUsecase) FetchImage(imageURL string) (*model.Image, error) {
	image := model.Image{}
	image.Filename = imageURL
	imageFile, err := iu.ir.FetchImage(&image)
	if err != nil {
		return nil, err
	}

	image.ImageFile = imageFile.ImageFile

	return &image, nil
}
