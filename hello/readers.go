package main

import (
  "fmt"
  "io"
  "strings"
)

func readers() {
  fmt.Println("* Readers *")

  // Many impl of Reader interface.
  // from io interface: populates given slice with data
  // func (T) Read(b []byte) (n int, e error)

  r := strings.NewReader("Hello, reader")
  b := make([]byte, 8)
  // Consume reader output 8 bytes at a time
  for {
    n, err := r.Read(b)
    fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
    fmt.Printf("b[:n] = %q\n", b[:n])
    if err == io.EOF {
      break
    }
  }

  // Exercises
  runReaders()
  runRot13Reader()
  fmt.Println()
}