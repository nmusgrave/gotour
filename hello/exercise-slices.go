package main

import (
  "golang.org/x/tour/pic"
  "fmt"
)

/*
 * Returns slice of length dy
 * Each element is slice of dx uint8
 * x ^ y
 */
func Pic(dx, dy int) [][]uint8 {
  pic := make([][]uint8, dy)
  for y := range pic {
    // Iterate over all rows
    row := make([]uint8, dx)
    for x := range row {
      row[x] = uint8((x + y)/2)
    }
    pic[y] = row
  }
  return pic
}

func runPic() {
  fmt.Println("* Slices Exercise *")
  pic.Show(Pic)
}