package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{0, 208, 35, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

// this will print green rectangular of 100x100 pixels
