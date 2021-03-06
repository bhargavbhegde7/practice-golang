package main

import(
  "fmt"
  "strconv"
  "strings"
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

func getInOrderString(root *Node) string{
  if root == nil{
    return ""
  }

  return strings.Trim(fmt.Sprint(
    getInOrderString(root.left) +" "+
    strconv.Itoa(root.data) +" "+
    getInOrderString(root.right)),
    " ")
}

func getPreOrderString(root *Node) string{
  if root == nil{
    return ""
  }

  return fmt.Sprint(
    strconv.Itoa(root.data)+
    getPreOrderString(root.left)+
    getPreOrderString(root.right))
}

func getPostOrderString(root *Node) string{
  if root == nil{
    return ""
  }

  return fmt.Sprint(
    getPostOrderString(root.left)+
    getPostOrderString(root.right)+
    strconv.Itoa(root.data))
}

func getHeight(root *Node) int{
  if root == nil{
    return 0
  }else{
    heightL := getHeight(root.left)
    heightR := getHeight(root.right)

    if heightL > heightR{
      return heightL + 1
    }else{
      return heightR + 1
    }
  }
}



func getLevelOrderString(root *Node){

}
