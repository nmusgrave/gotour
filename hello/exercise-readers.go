package main

import (
  "fmt"
  "golang.org/x/tour/reader"
)

// emits infinite stream of 'A' characters
type MyReader struct {}

func (r MyReader) Read(b []byte) (n int, e error) {
  for i := range b {
    b[i] = 'A'
  }
  return len(b), nil
}

func runReaders() {
  fmt.Println("* Reader Exercise *")
  reader.Validate(MyReader{})
}
