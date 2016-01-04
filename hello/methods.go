package main

import (
  "fmt"
  "math"
)

var f = MyFloat(-math.Sqrt2)
var v = &Vertex{3, 4}

// Method receiver: arglist between 'func' and name of function
// Assoc w pointer to named type
func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Assoc w named type
func (m MyFloat) Abs() float64 {
  if m < 0 {
    return float64(-m)
  }
  return float64(m)
}

func methods() {
  fmt.Println("-- Basics of Methods --")

  // No classes. Methods on struct types are next best
  v.Abs()

  // Can have method on any type declared in package, but not from other pkgs,
  // or built in types
  fmt.Println(f.Abs())

  // Can be assoc w named type, or ptr to named type
  // Prefer ptr: doesn't copy entire value on invocation, and can modify value
  // referenced by ptr

  interfaces()
  errors()
  readers()
  servers()
  images()
}

type MyFloat float64
