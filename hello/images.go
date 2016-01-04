package main

import (
  "fmt"
  "image"
)

func images() {
  fmt.Println("* Images *")
  fmt.Println("Image pkg defines interface ColorModel(), Bounds(), At(int,int)")

  m := image.NewRGBA(image.Rect(0,0,100,100))
  fmt.Println(m.Bounds())
  fmt.Println(m.At(0,0).RGBA())

  // color.Color and color.Model are interfaces.
  // Can use built in impl color.RGBA and color.RGBAModel (as in image.color pkg)

  // Exercise
  runImages()
}