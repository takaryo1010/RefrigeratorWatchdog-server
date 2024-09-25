package model

import (
	"io"
)

type Image struct {
	ImageFile io.Reader
	Filename  string
}

