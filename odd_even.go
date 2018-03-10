package main

import(
  "fmt"
)

func oddProducer(oddChannel chan int, f func(string)){
  for i := 0; i<=100; i++{
    if i%2 != 0 {
      oddChannel<-i
    }
  }
  oddChannel <- -1
  f("odd producer finished")
}

func evenProducer(evenChannel chan int, f func(string)){
  for i := 0; i<=100; i++{
    if i%2 == 0 {
      evenChannel<-i
    }
  }
  evenChannel <- -1
  f("even producer finished")
}

func consumer(oddChannel chan int, evenChannel chan int, f func(string)){
  oddDone := false
  evenDone := false

  for {
    select{
      case oddNum := <-oddChannel:
        if oddNum < 0 {
          oddDone = true
        }
        fmt.Println("odd  : ",oddNum)
      case evenNum := <-evenChannel:
        if evenNum < 0 {
          evenDone = true
        }
        fmt.Println("even : ",evenNum)
    }
    if oddDone && evenDone {
      break
    }
  }
  f("consumer finished")
}

func main(){
  oddChannel := make(chan int)
  evenChannel := make(chan int)

  onFinished := func(message string) {
         fmt.Println(message)
      }

  go oddProducer(oddChannel, onFinished);
  go evenProducer(evenChannel, onFinished);
  go consumer(oddChannel, evenChannel, onFinished);

  var input string
  fmt.Scanln(&input)
}
