package main

import (
  "fmt"
  "math"
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
  printSlice("a", a)

  capacity := length + 1
  b := make([]int, length, capacity)
  printSlice("b", b)

  // length of 3, same capacity as original
  c := s[:3]
  printSlice("c", c)

  // re-grow length of slice to its original capacity
  c = c[:cap(c)]
  printSlice("c", c)

  fmt.Println("Increasing the slice size: make new and copy. Using copy() currectly handles slices w same underlying array")
  printSlice("s", s)
  t := make([]int, len(s), 2 * (cap(s) + 1))
  for i := range s {
    t[i] = s[i]
  }
  s = t
  printSlice("s", s)
  // Simpler way
  t = make([]int, 2 * (len(s) + 1))
  copied := copy(t, s)
  printSlice("t", t)
  fmt.Printf("Moved %d elements\n", copied)
}

func appendingSlices() {
  fmt.Println("* Appending Slices *")
  fmt.Println("Appending to a slice, using custom")
  t := []int{1, 3, 5, 7}
  printSlice("Original", t)
  t = appendInt(t, 9, 11)
  printSlice("Append1", t)
  t = appendInt(t, 9, 11)
  printSlice("Append2", t)
  t = appendInt(t, 13)
  printSlice("Append3", t)

  fmt.Println("Appending to a slice, using BIF. Note capacity only grown when needed.")
  t = []int{1, 3, 5, 7}
  printSlice("Original", t)
  t = append(t, 9, 11)
  printSlice("Append1", t)
  t = append(t, 9, 11)
  printSlice("Append2", t)
  t = append(t, 13)
  printSlice("Append3", t)

  fmt.Println("Combinining slices using append")
  t = append(t, t...) // equivalent to "append(t, t[0], t[1], t[2])"
  printSlice("Combined", t)
}

// custum impl. go has built in support
func appendInt(slice []int, data ...int) []int {
  sliceLength := len(slice)
  capacity := sliceLength + len(data)
  // realloc if need more space
  if capacity > cap(slice) {
    bigger := make([]int, (capacity+1) * 2)
    copy(bigger, slice)
    slice = bigger
  }
  slice = slice[0:capacity]
  copy(slice[sliceLength:capacity], data)
  return slice
}

func printSlice(name string, s []int) {
  fmt.Println(name, s, len(s), cap(s))
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

var m map[string]Vertex

func maps() {
  fmt.Println("* Maps *")
  // Can init with make
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{40.68433, -74.39967}
  // or as a literal
  m = map[string]Vertex {
    "Bell Labs": Vertex{40.68433, -74.39967},
    "Google": Vertex{37.42202, -122.08408},
  }
  // can omit type names in elements if declared at top level
  m = map[string]Vertex {
    "Bell Labs": {40.68433, -74.39967},
    "Google": {37.42202, -122.08408},
  }
  fmt.Println(m)

  delete(m, "Bell Labs")
  elem, ok := m["Bell Labs"]
  fmt.Println("after delete, elem is zero-value, ok is false", elem, ok)
}

// Inputs: function that accepts two float64's, and returns a float64
// Returns: a float64
func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

// returns a closure (function that references variable outside its body)
// each closure has its own associated 'sum' variable
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func functions() {
  fmt.Println("* Functions *")
  fmt.Println("Functions can be passed around as values")

  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }

  fmt.Println(hypot(5, 7))
  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))

  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(pos(i), neg(-i))
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