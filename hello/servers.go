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
  // go to this with browser to see result

  // NOTE uncomment below to launch servers
  //log.Fatal(http.ListenAndServe("localhost:8000", h))
  //runHttpHandlers()
  fmt.Println(log.Ldate, h)
}