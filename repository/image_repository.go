package repository

import (
	"RefrigeratorWatchdog-server/model"
	"errors"
	"fmt"
	"io"
	"os"
)

type IImageRepository interface {
	UploadImage(image *model.Image) (*model.Image, error)
	FetchImage(image *model.Image) (*model.Image, error)
}

type imageRepository struct {
}

func NewImageRepository() IImageRepository {
	return &imageRepository{}
}

func (ir *imageRepository) UploadImage(file *model.Image) (*model.Image, error) {

	dst, err := os.Create("images/" + file.Filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file.ImageFile); err != nil {
		return file, err
	}

	return file, nil
}

func (ir *imageRepository) FetchImage(file *model.Image) (*model.Image, error) {
	if _, err := os.Stat("images/" + file.Filename); os.IsNotExist(err) {
		return nil, errors.New("ファイルが存在しません")
	}

	src, err := os.Open("images/" + file.Filename)
	if err != nil {
		return nil, fmt.Errorf("ファイルを開くことができません: %v", err)
	}

	_, err = src.Stat()
	if err != nil {
		return nil, fmt.Errorf("ファイル情報の取得に失敗: %v", err)
	}
	file.ImageFile = src

	return file, nil
}
