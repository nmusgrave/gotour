package main

import (
  "fmt"
  "time"
)

func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func sum(a []int, ch chan int) {
  sum := 0
  for _, v := range a {
    sum += v
  }
  // send
  ch <- sum
}

func fib(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  // Only sender closes the channel
  close(c)
  // sending on closed channel gives panic
  //c <- 0
}

func fibSelect(c chan int, quit chan int) {
  x, y := 0, 1
  for {
    select {
    case c <- x:
      x, y = y, x+y
    case <-quit:
      fmt.Println("quit")
      return
    }
  }
}

func channels() {
  fmt.Println("* Channels *")
  fmt.Println("Typed way to send & receive msgs. Data flows in direction of arrow.")
  fmt.Println("Default: send & receive block till other side ready. Provides synchronization")

  a := []int{7, 2, 8, -9, 0, 5}
  // init
  ch := make(chan int)

  go sum(a[:len(a)/2], ch)
  go sum(a[len(a)/2:], ch)
  // receive
  x, y := <-ch, <-ch
  fmt.Println("Sum:", x, y, x + y)

  fmt.Println("Buffered channels: provide length when init. Sends block when full, receives block when empty")
  ch = make(chan int, 2)
  ch <- 1
  ch <- 2
  fmt.Println(<-ch, <-ch)

  fmt.Println("Sender can close channel, and reciever can check if closed. Don't have to close.")
  fmt.Println("")
  // Receiver determines if channel open
  ch = make(chan int, 10)
  go fib(cap(ch), ch)
  // alt: can iterate over values using range. Stops when channel closed
  for {
    i, ok := <-ch
    fmt.Println(i)
    if ok {
      fmt.Println("channel open")
    } else {
      fmt.Println("channel closed by sender")
      break
    }
  }

  fmt.Println("select: goroutine waits on multiple connections. Blocks till a case can run. Chooses randomly if multiple ready.")
  quit := make(chan int)
  ch = make(chan int)
  go func() {
    for i := 0; i < 10; i++ {
      fmt.Println(<-ch)
    }
    quit <- 0
  }()
  fibSelect(ch, quit)

  fmt.Println("Can use default select case to send/receive w/o blocking. Run when no other case ready")
  tick := time.Tick(100 * time.Millisecond)
  boom := time.After(200 * time.Millisecond)
  for {
    select {
    case <-tick:
      fmt.Println("tick")
    case <-boom:
      fmt.Println("boom")
      return
    default:
      // receiving from tick or boom would block
      fmt.Println("    .")
      time.Sleep(50 * time.Millisecond)
    }
  }
}

func concurrency() {
  fmt.Println("-- Basics of Concurrency --")
  fmt.Println("goroutine: lightweight thread. Evaluation of func and args in current routine, execution in new.")
  fmt.Println("All in same addr space, so must sync access to shared mem")

  // Start new go routine running say("baby").
  go say("baby")
  say("hello")

  channels()
}