package main

import (
  "fmt"
)

// Basic dereferencing
func pointers() {
  fmt.Println("* Pointers *")
  i := 7
  var p *int
  p = &i
  fmt.Println(p, *p == i)
  i = 23
  fmt.Println(p, *p == i)
  *p = 46
  fmt.Println(p, *p == i)
}

// 'struct' type containing named fields
// 'type' introducces new type
type Vertex struct {
  X, Y float64
}

func structInit() {
  // zero-init fields
  var v1 *Vertex = new(Vertex)
  // or do this to get type *Vertex
  v1 = &Vertex{1, 2}
  // Explicit init of fields
  var v2 Vertex = Vertex{X: 1, Y: 1}
  // or can use this if know order of fields
  v2 = Vertex{2, 2}
  // or init subset of fields
  v2 = Vertex{Y: 3}
  fmt.Println(*v1, v2)
}

func ranges() {
  fmt.Println("Note: generated slice points to array of original contents. If want to release memory, need to copy out elements")
  fmt.Println("Range over slice returns index and copy of element")
  t := []int{1, 3, 5}
  for i, e := range t {
    fmt.Println(i, e)
  }
  // Can drop index
  for _, e := range t {
    fmt.Println(e)
  }
  // or drop value entirely
  for i := range t {
    t[i] += 1
    fmt.Println(t[i])
  }
}

func types() {
  fmt.Println("-- Basics of Types --")
  
  // zero-value
  var z []int
  fmt.Println("empty slice is nil", z == nil)

  // No pointer arithmetic
  //p += 1

  pointers()
  structInit()
  arraysVsSlices()
  buildingSlices()
  appendingSlices()
  ranges()
  maps()
  functions()

  // Exercises for data structures
  runPic()
  runWordCount()
  // display the first 10 fib nums
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}