package main

import (
  "fmt"
  "gotour/stringutil"
  "math"
  "math/rand"
)


func main() {
  fmt.Printf("Hello\n")
  fmt.Printf(stringutil.Reverse("Hello") + "\n")
  fmt.Println("Some num", rand.Intn(10), math.Pi)
  basics()
  flow()
}