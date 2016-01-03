package main

import (
  "time"
  "fmt"
  "strconv"
)

type MyError struct {
  When time.Time
  What string
}

// MyError implicitly implements Error interface
func (e *MyError) Error() string {
  return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func runError() error {
  return &MyError{time.Now(), "didn't work"}
}

func errors() {
  fmt.Println("* Errors *")

  // Built-in interface, like stringers. Expresses error values.
  i, err := strconv.Atoi("42")
  if err != nil {
    // some failure found
    fmt.Printf("couldn't convert number %v\n", err)
  } else {
    fmt.Printf("converted %v\n", i)
  }

  fmt.Println(runError())

  // Exercises
  runErrors()
}