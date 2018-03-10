package main

import(
  "fmt"
)

func oddProducer(oddChannel chan int){
  for i := 1; i<100; i++{
    if i%2 != 0 {
      oddChannel<-i
    }
  }
}

func evenProducer(evenChannel chan int){
  for i := 0; i<100; i++{
    if i%2 == 0 {
      evenChannel<-i
    }
  }
}

func consumer(oddChannel chan int, evenChannel chan int){
  for i := 0; i<100; i++{
    select{
      case oddNum := <-oddChannel:
        fmt.Println("odd  : ",oddNum)
      case evenNum := <-evenChannel:
        fmt.Println("even : ",evenNum)
    }
  }
}

func main(){
  oddChannel := make(chan int)
  evenChannel := make(chan int)

  go oddProducer(oddChannel);
  go evenProducer(evenChannel);
  go consumer(oddChannel, evenChannel);

  var input string
  fmt.Scanln(&input)
}
