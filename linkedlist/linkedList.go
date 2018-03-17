package main

import(
  "fmt"
  "strings"
)

type Node struct{
  data int
  next *Node
}

type LinkedList struct{
  head *Node
}

func (l *LinkedList) AddToHead(data int){
  l.head = &Node{data: data, next: l.head}
}

func getLastNode(head *Node) *Node{
  cur := head
  for{
    if cur.next == nil{
      break
    }else{
      cur = cur.next
    }
  }
  return cur
}

func addToTail(head *Node, data int){
  lastNode := getLastNode(head)
  lastNode.next = &Node{data: data, next: nil}
}

func printAll(head *Node) string{
  output := ""
  cur := head
  for{
    if cur.next == nil{
      break
    }else{
      output += fmt.Sprintf("%d%s", cur.data, " ")
      cur = cur.next
    }
  }

  return strings.Trim(output, " ")
}

func getExampleList() *Node{
  head := &Node{data: 0, next: nil}
  for i:= 1; i <= 10; i++{
    addToTail(head, i*10)
  }
  return head
}

func main(){
  head := getExampleList()
  if strings.EqualFold(printAll(head), "0 10 20 30 40 50 60 70 80 90"){
    fmt.Println("success")
  }
}
