package controller

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/usecase"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IImageController interface {
	UploadImage(c echo.Context) error
	FetchImage(c echo.Context) error
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

	defer func() {
		if closer, ok := file.ImageFile.(io.Closer); ok {
			closer.Close()
		}
	}()

	image, err := ic.iu.UploadImage(file)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, image)
}

func (ic *imageController) FetchImage(c echo.Context) error {
	imageURL := c.Param("imageURL")
	image, err := ic.iu.FetchImage(imageURL)
	if err != nil {
		return c.JSON(400, err)
	}

	fileHeader := make([]byte, 512)
	image.ImageFile.Read(fileHeader)
	if seeker, ok := image.ImageFile.(io.Seeker); ok {
		seeker.Seek(0, 0)
	}
	mimeType := http.DetectContentType(fileHeader)

	defer func() {
		if closer, ok := image.ImageFile.(io.Closer); ok {
			closer.Close()
		}
	}()

	return c.Stream(http.StatusOK, mimeType, image.ImageFile)
}
