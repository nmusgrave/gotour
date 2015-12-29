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
  X, Y int
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

func arraysVsSlices() {
  fmt.Println("* Arrays vs Slices *")
  // arrays store values
  // represents the entire structure, rather than just ptr to structure
  greeting := "hey"
  var a [2]string
  a[0] = greeting
  a[1] = "baby"
  fmt.Println("array", a)
  greeting = "bonjour"
  fmt.Println("Modifying variable not shown in array contents. a[1] same", a)

  // array literal
  a = [...]string{"hello", "dolly"}

  // slice: ptr to array, length, capacity. Can have any type, including other slices
  i := 5
  // Slice literal has no element count
  s := []int{0, 1, i, 8, 9, 32}
  fmt.Println("slice", s)
  i += 1
  fmt.Println("Modifying variable not shown in slice contents. s[2] same", s)

  // Re-slicing generates new slice that references previous slice
  lo := 0 // inclusive
  hi := 3 // exclusive
  fmt.Println("empty slice", len(s[lo:lo]) == 0)
  fmt.Println("one-element slice", len(s[lo:lo+1]) == 1)
  short := s[lo:hi]
  fmt.Println("longer slice", len(short) > 1)
  short[0] = 100
  fmt.Println("Modifying new slice changes previous slice", s[0] == short[0])
}

func buildingSlices() {
  fmt.Println("* Building Slices *")
  s := []int{0, 1, 5, 8, 9, 32}
  // make() allocs zeroed array, returns slice referring to that array
  length := 5
  a := make([]int, length)
  printSlice(a)

  capacity := length + 1
  b := make([]int, length, capacity)
  printSlice(b)

  // length of 3, same capacity as original
  c := s[:3]
  printSlice(c)
}

func printSlice(s []int) {
  fmt.Println(s, len(s), cap(s))
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

}