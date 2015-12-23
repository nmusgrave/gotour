package main

import (
  "fmt"
  "unicode/utf8"
)

// Variable declaration: type comes after variable name
// Can factor variable declaraion
var (
  p *int
  a [3]int
  f func(func(int, int) int, int) int
)

// Constants use 'const' rather than 'var' keyword
const (
  PI = 3.14
  greeting = "bonjour"
)

// 'var' declares list of variables
// Can init and leave out type
var j, k = false, true

// Can share type if all same 
func add(a, b int) int {
  return a + b
}

func mult(f func(int, int) int, n int) int {
  return f(n, n)
}

// Return multiple results at once
func swap(a, b string) (string, string) {
  return b, a
}

// Named return: return variables have names
func swapNaked(a, b string) (s1, s2 string) {
  s1 = b
  s2 = a
  // Naked return: empty return. Works when have named variables
  return
}

func basics() {
  fmt.Println("-- Exploring Basics of Go Syntax --")
  
  // Testing strange behavior of range for utf-8 encodings
  valid := []byte("Hello, 世界")
  fmt.Println(utf8.Valid(valid))
  for len(valid) > 0 {
    r, size := utf8.DecodeRune(valid)
    fmt.Printf("%c %v\n", r, size)
    valid = valid[size:]
  }

  invalid := []byte{0xff, 0xfe, 0xfd}
  s := string(invalid)
  fmt.Println(utf8.Valid(invalid))
  for index, r := range s {
    //r, size := utf8.DecodeRune(valid)
    fmt.Printf("%c %v\n", r, index)
  }

  // Function pointers
  fmt.Println(mult(add, 5))
  f = mult
  fmt.Println(f(add, 5))

  // Return multiple results
  // := short assignment used in place of 'var', handles types implicitly. For variables inside functions
  x, y := swap("bae", "hey")
  fmt.Println(x, y)

  // Named, naked return
  x, y = swapNaked("soir", "bonne")
  fmt.Println(x, y)

  // Constants don't use short assignment
  const word = "word"
  fmt.Println(greeting)
}