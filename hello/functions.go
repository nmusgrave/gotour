package main

import (
  "fmt"
  "math"
)

// Inputs: function that accepts two float64's, and returns a float64
// Returns: a float64
func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

// returns a closure (function that references variable outside its body)
// each closure has its own associated 'sum' variable
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func functions() {
  fmt.Println("* Functions *")
  fmt.Println("Functions can be passed around as values")

  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }

  fmt.Println(hypot(5, 7))
  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))

  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(pos(i), neg(-i))
  }
}