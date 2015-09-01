/**
Exercise: Images
Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an
implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
 */
package main

import (
	"image"
	"letsgo/pic"
	"image/color"
)

type Image struct{
	Width, Height int
	colr uint8
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colr+uint8(x), r.colr+uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
