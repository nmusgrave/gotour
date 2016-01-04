package main

import (
  "fmt"
  "golang.org/x/tour/pic"
  "image/color"
  "image"
)

type Image struct{}

func (m Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, 256, 256)
}

func (m Image) At(x, y int) color.Color {
  v := uint8((x + y)/2)
  return color.RGBA{v, v, 255, 255}
}

func runImages() {
  fmt.Println("* Images Exercise *")
  m := Image{}
  pic.ShowImage(&m)
}