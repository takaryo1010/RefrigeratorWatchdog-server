package model

import "mime/multipart"

type Image struct {
	ImageFile multipart.File
	Filename string
}
