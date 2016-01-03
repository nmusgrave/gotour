package main

import (
  "os"
  "io"
  "strings"
  "regexp"
  "fmt"
)

// Wrap reader, and modify it
// Modifies stream by applying rot13 substitution cipher to all alphabetical chars
// rot13: rotate by 13 places
type rot13Reader struct {
  r io.Reader
}

// Implements Reader interface
func (reader rot13Reader) Read(b []byte) (n int, err error) {
  n = 0
  b2 := make([]byte, 8)
  for {
    // Read in bytes
    num, e := reader.r.Read(b2)
    if e == io.EOF {
      return n, io.EOF
    }
    // Rotate alphabetical bytes
    for i, elem := range b2 {
      r, e := regexp.Compile(`[[:alpha:]]`)
      if e == nil && r.MatchString(string(elem)) == true {
        var adjust uint8 = 'A'
        if elem >= 'a' {
          adjust = 'a'
        }
        elem = (elem - adjust + 13) % 26 + adjust
      }
      b2[i] = byte(elem)
    }
    // Save adjusted bytes to output string
    copy(b[n:], b2)
    n += num
  }
  return n, nil
}

func runRot13Reader() {
  fmt.Println("* rot13 Exercise *")
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}