package main

import (
  "fmt"
  "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  // must first convert to float64, to avoid recursively calling Error() till out of memory
  return fmt.Sprintf("cannot Sqrt negative num %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
  // Approx w newton's method: z = z - (z^2 - x)/2z
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }
  z := 1.0
  for newz := z - (z * z - x) / (2 * z); math.Abs(z - newz) > 0.001; newz = z - (z * z - x) / (2 * z) {
    z = newz
  }
  return z, nil
}

func runErrors() {
  fmt.Println("* Errors Exercise *")
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}