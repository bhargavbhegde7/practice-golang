package main

import(
  "fmt"
)

func main(){
  oddChannel := make(chan int)

  go func(){
    i := <-oddChannel
    fmt.Println("HJELEL", i)
    }()

  oddChannel <- 42

  var input string
  fmt.Scanln(&input)
}
