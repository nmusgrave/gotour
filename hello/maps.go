package main

import (
  "fmt"
)

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
