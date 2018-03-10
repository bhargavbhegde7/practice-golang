package main

import (
  "fmt"
)

func producer(c chan int) {
  for i := 0; i<100 ; i++ {
    c <- i
  }
  //close(c)
}

func oddConsumer(c chan int) {
  for {
    msg := <- c

    if msg % 2 != 0{
      fmt.Println("odd  : ", msg)
      //time.Sleep(time.Second * 1)
    }else{
      c <- msg
    }
  }
}

func evenConsumer(c chan int) {
  for {
    msg := <- c

    if msg % 2 == 0{
      fmt.Println("even : ", msg)
      //time.Sleep(time.Second * 1)
    }else{
      c <- msg
    }
  }
}

func main() {
  c := make(chan int)

  go producer(c)
  go oddConsumer(c)
  go evenConsumer(c)

  var input string
  fmt.Scanln(&input)
}
