package main

import (
  "fmt"
  "math"
  "os"
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
}

type MyFloat float64
