package main

import (
  "fmt"
  "sync"
  "time"
)

type SafeCounter struct {
  v map[string]int
  mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
  c.mux.Lock()
  c.v[key]++
  c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
  c.mux.Lock()
  defer c.mux.Unlock()
  return c.v[key]
}

func mutex() {
  fmt.Println("* Mutex *")
  fmt.Println("Sync w/o communication, using locking")
  fmt.Println("defer: ensure mutex unlocked")

  c := SafeCounter{v: make(map[string]int)}
  for i := 0; i < 1000; i++ {
    go c.Inc("somekey")
  }

  time.Sleep(time.Second)
  fmt.Println(1000 == c.Value("somekey"))

  // Exercise
  runWebCrawler()
}