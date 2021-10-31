package main

import (
	"image"
	"image/color"

	"github.com/golang-id/tour/pic"
)

type Image struct {
	rows, cols int
}

func (m *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.rows, m.cols)
}

func (m *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{64, 64}
	pic.ShowImage(&m)
}
