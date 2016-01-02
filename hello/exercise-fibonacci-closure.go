package main

// returns function that returns an int
func fibonacci() func() int {
  fib0 := 0
  fib1 := 1
  return func() int {
    result := fib0
    next := fib0 + fib1
    fib0 = fib1
    fib1 = next
    return result
  }
}