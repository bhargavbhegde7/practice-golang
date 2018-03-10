package main

import(
  "fmt"
  "time"
)

func main() {
  c1 := make(chan int)
  c2 := make(chan int)

  go func() {
    i1 := 0
    for {
      c1 <- i1
      i1 = i1 + 2
      time.Sleep(time.Second * 2)
    }
  }()

  go func() {
    i2 := 1
    for {
      c2 <- i2
      i2 = i2 + 2
      time.Sleep(time.Second * 3)
    }
  }()

  go func() {
    for {
      select {
      case msg1 := <- c1:
        fmt.Println("even : ", msg1)
      case msg2 := <- c2:
        fmt.Println("odd  : ", msg2)
      }
    }
  }()

  var input string
  fmt.Scanln(&input)
}
