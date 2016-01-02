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
    fmt.Printf("%f >= %g\n", v, lim)
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

func run() {
  defer func() {
    if r := recover(); r != nil {
      // This will catch the panic, and restore normal execution
      // By putting this inside a defer, ensures that any panic from any function
      //  in body will be caught
      fmt.Println("Recover in run()", r)
    }
  }()
  count(0)
  fmt.Println("Returned normally from count()")
}

func count(i int) {
  if i > 3 {
    fmt.Println("Panicking", i)
    panic(fmt.Sprintf("%v", i))
  }
  // These will execute in reverse order, as unwind call stack from panic
  defer fmt.Println("defer in count()", i)
  // These will execute in increasing order, before panic
  fmt.Println("count() of ", i)
  count(i + 1)
}

func flow() {
  fmt.Println("-- Basics of Flow Control --")

  // infinite loop
  //for {}

  // Arguments to defered function evaluated early, function invoked when surrounding function returns
  // 'defer' pushes function calls onto stack
  // Allowed to modify named return variables
  // Usually for cleanup operations
  defer fmt.Println("defered greeting:", giveGreeting())

  x := 5
  n := 2
  fmt.Printf("%d summed %d times: %d\n", x, n, summer(x, n))
  n = 31
  fmt.Printf("sum %d while result < %d: %d\n", x, n, boundedSummer(x, n))

  base := 2.0
  exp := 3.0
  lim := 7.0
  fmt.Printf("%g ^ %g = %g if under limit %g\n", base, exp, pow(base, exp, lim), lim)
  fmt.Println("sqrt(70) is 8.36660:", sqrt(70))

  identifyOS()

  // Demonstration of handling defer, panic, recover
  run()
  fmt.Println("Returned from run()")

}