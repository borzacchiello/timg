package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/gabriel-vasile/mimetype"
)

func LoadImage(data []uint8) (image.Image, error) {
	mime := mimetype.Detect(data)
	switch mime.String() {
	case "image/png":
		return png.Decode(bytes.NewReader(data))
	case "image/jpeg":
		return jpeg.Decode(bytes.NewReader(data))
	}
	return nil, fmt.Errorf("unsupported image type '%s'", mime.String())
}
