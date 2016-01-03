package main

import (
  "fmt"
  "log"
  "net/http"
)

type Hello struct {}

// Implements the Handler interface
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "hello")
}

func servers() {
  fmt.Println("* Servers *")
  fmt.Println("HTTP package serves requests using any value impl Handler")

  var h Hello
  err := http.ListenAndServe("localhost:4000", h)
  if err := nil {
    log.Fatal(err)
  }
}