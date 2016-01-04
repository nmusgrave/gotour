package main

import (
  "fmt"
  "log"
  "net/http"
)

type String string

type Struct struct {
  Greeting string
  Punct string
  Who string
}

// Implements the Handler interface
func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, h)
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, h)
}

func runHttpHandlers() {
  fmt.Println("* Exercise - http handlers *")
  fmt.Println("HTTP package serves requests using any value impl Handler")

  // go to /string and /struct with browser to see result
  http.Handle("/string", String("I'm a frayed knot"))
  http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
  http.Handle("/", String("Hello"))

  tcpAddr := "localhost:4000"
  log.Fatal(http.ListenAndServe(tcpAddr, nil))
}