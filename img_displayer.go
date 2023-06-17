package main

import (
	"fmt"
	"image"
	"image/color"

	tsize "github.com/kopoli/go-terminal-size"
	"github.com/nfnt/resize"
)

// const gradient = "@%#*+=-:. "
const gradient = " .:-=+*#%@"

func getCharFromPixel(c color.Color) byte {
	r, g, b, a := c.RGBA()
	intensity := ((float32(r) + float32(g) + float32(b)) / (3 * 0xffff)) * (float32(a) / 0xffff)
	idx := int(intensity*float32(len(gradient)-1) + 0.5)
	return gradient[idx]
}

func DisplayImage(img image.Image) error {
	s, err := tsize.GetSize()
	if err != nil {
		return err
	}

	newImg := resize.Resize(0, uint(s.Height-2), img, resize.Lanczos3)
	if newImg.Bounds().Dx() > s.Width/2 {
		newImg = resize.Resize(uint(s.Width/2), 0, img, resize.Lanczos3)
	}

	for y := 0; y < newImg.Bounds().Dy(); y++ {
		for x := 0; x < newImg.Bounds().Dx(); x++ {
			c := getCharFromPixel(newImg.At(x, y))
			fmt.Printf("%c%c", c, c)
		}
		fmt.Printf("\n")
	}
	return nil
}
