package main

import(
  "fmt"
  "time"
)

func main(){
 channel := make(chan string)

 go producer(channel, "odd");
 go producer(channel, "even");

 channel<-"odd"

 var input string
 fmt.Scanln(&input)
}

func producer(channel chan string, msg string){
  for {
    time.Sleep(time.Second)
    fmt.Println(<-channel)
    channel<-msg
  }
}
