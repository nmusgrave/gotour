package main

import (
  "fmt"
)

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