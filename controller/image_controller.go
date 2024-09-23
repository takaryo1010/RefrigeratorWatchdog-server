package controller

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/usecase"

	"github.com/labstack/echo/v4"
)

type IImageController interface {
	UploadImage(c echo.Context) error
	
}

type imageController struct {
	iu usecase.IImageUsecase
}

func NewImageController(iu usecase.IImageUsecase) IImageController {
	return &imageController{iu}
}

func (ic *imageController) UploadImage(c echo.Context) error {
	imageFile, err := c.FormFile("image")
	if err != nil {
		return c.JSON(400, err)
	}
	var file model.Image
	file.ImageFile, err = imageFile.Open()
	if err != nil {
		return c.JSON(400, err)
	}
	file.Filename = imageFile.Filename

	defer file.ImageFile.Close()

	image, err := ic.iu.UploadImage(file)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, image)
}
