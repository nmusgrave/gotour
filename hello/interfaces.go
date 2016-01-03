package main

import (
  "fmt"
  "os"
)

// Interface type defined by set of methods.
// Type implements interface by implementing methods.
type Abser interface {
  Abs() float64
}

// No explicit keywords for implementing some package
// Implicit interface impl decouple impl packages and declarations
type Reader interface {
  Read(b []byte) (n int, err error)
}

type Writer interface {
  Write(b []byte) (n int, err error)
}

type ReaderWriter interface {
  Reader
  Writer
}

// Implicitly implements Stinger interface
type Person struct {
  Name string
  Age int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func interfaces() {
  fmt.Println("* Interfaces *")

  var a Abser

  // MyFloat impl Abser
  a = f
  a.Abs()
  // *Vertex impl Abser
  a = v
  a.Abs()
  // 'a' is Vertex, doesn't impl Abser. Assignment invalid
  //a = *v

  // Implicit Interfaces
  var w Writer
  w = os.Stdout // implements Writer
  // Uses declared writer to output to console
  fmt.Fprintf(w, "hello, writer\n")

  // Stringer interface (def'd in 'fmt' pkg) impl by any type that can define
  // itself as a string.
  me := Person{"Naomi", 21}
  fmt.Println(me)

  runStringer()
}