package main

import(
  "fmt"
  //"strings"
)

type Node struct{
  data int
  left *Node
  right *Node
}

func insert(val int, root *Node){
  cur := root
  for{
    if val < cur.data {
      if cur.left != nil{
        cur = cur.left
      }else{
        cur.left = &Node{data: val, left: nil, right: nil}
        break
      }
    }else if val > cur.data{
      if cur.right != nil{
        cur = cur.right
      }else{
        cur.right = &Node{data: val, left: nil, right: nil}
        break
      }
    }else{
      break
    }
  }//for ends
}

/**
     *
     * @return        8
     *              /   \
     *             /     \
     *            3        10
     *           / \        \
     *          1   6        14
     *             / \       /
     *            4   7    13
     */
func getExampleBST() *Node{
  root := &Node{data: 8, left: nil, right: nil}

  insert(3, root)
  insert(10, root)
  insert(1, root)
  insert(6, root)
  insert(14, root)
  insert(4, root)
  insert(7, root)
  insert(13, root)

  return root
}

func main(){
  root := getExampleBST()

  fmt.Println(root.data)
  fmt.Println(root.left.data)
  fmt.Println(root.right.data)
}
