package main

import (
  "fmt"
  "golang.org/x/tour/tree"
)

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

// Walk the tree, sending all values into channel
func Walk(t *tree.Tree, ch chan int) {
  WalkHelper(t, ch)
  close(ch)
}

func WalkHelper(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  }
  WalkHelper(t.Left, ch)
  ch <- t.Value
  WalkHelper(t.Right, ch)
}

// use channels & concurrency to determine if two trees are the same
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  go Walk(t1, ch1)
  ch2 := make(chan int)
  go Walk(t2, ch2)

  for i := range ch1 {
    j := <- ch2
    if i != j {
      return false
    }
  }
  return true
}

func runEquivTrees() {
  fmt.Println("* Equivalent Binary Trees Exercise *")

  // test walk
  ch := make(chan int, 10)
  go Walk(tree.New(1), ch)
  for i := 0; i < cap(ch); i++ {
    fmt.Println(<-ch)
  }

  // test same
  fmt.Println(true == Same(tree.New(1), tree.New(1)))
  fmt.Println(false == Same(tree.New(1), tree.New(2)))
}