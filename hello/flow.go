package main

import (
  "fmt"
  "math"
  "runtime"
)

// basic 'for' loop
func summer(a, n int) (sum int) {
  sum = 0
  for i := 0; i < n; i++ {
    sum += a
  }
  return
}

// 'while' can be expressed via 'for' loop
func boundedSummer(a, n int) (sum int) {
  sum = 0
  for sum < n {
    sum += a
  }
  return
}

// 'if' with short statement
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    // 'v' can be used in subsequent conditionals, too
    fmt.Printf("%g >= %g\n", v, lim)
  }
  return lim
}

func sqrt(x float64) float64 {
  // Approx w newton's method: z = z - (z^2 - x)/2z
  z := 1.0
  for newz := z - (z * z - x) / (2 * z); math.Abs(z - newz) > 0.001; newz = z - (z * z - x) / (2 * z) {
    z = newz
  }
  return z
}

// Can also write without switch condition
func identifyOS() {
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X")
  case "linux":
    fmt.Println("Linux")
  default:
    fmt.Println(os)
  }
}

func giveGreeting() string {
  fmt.Println("eval giveGreeting()")
  return "hi"
}

func flow() {
  fmt.Println("-- Basics of Flow Control --")

  // Arguments to defered function evaluated early, function invoked when surrounding function returns
  defer fmt.Println("Greeting:", giveGreeting())

  fmt.Println(summer(5, 2))
  fmt.Println(boundedSummer(5, 31))

  // infinite loop
  //for {}

  fmt.Println(pow(2, 3, 7))
  fmt.Println(sqrt(70))

  identifyOS()

}