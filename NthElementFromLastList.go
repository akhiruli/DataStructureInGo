package main

import "fmt"

type Node struct{
   next  *Node
   key   interface{}
}

type List struct{
   head  *Node
}

func Insert(root *Node, val interface{}) *Node{
   newNode := &Node{
      key: val,
      next: nil,
   }

   if root == nil {
      root = newNode
   }else{
      newNode.next = root
      root = newNode
   }

   return root
}

/*
1) Calculate the length of Linked List. Let the length be len.
2) Print the (len – n + 1)th node from the begining of the Linked List.
*/

func printNthFromLast(root *Node, n int){
   tmp := root
   length := 0
   for tmp != nil {
      length = length + 1
      tmp = tmp.next
   }

   if length < n {
      fmt.Println("Length of linkedlist is less")
      return
   }

   tmp = root

   for i := 1; i < length - n + 1; i++{
      tmp = tmp.next
   }

   fmt.Printf("%dth element from end: %v\n", n, tmp.key)
}

func Display(root *Node) {
   for root != nil {
      fmt.Printf("%+v ->", root.key)
      root = root.next
   }

   fmt.Println()
}

func main(){
   var root *Node
   root = Insert(root, 10)
   root = Insert(root, 20)
   root = Insert(root, 30)
   root = Insert(root, 40)
   root = Insert(root, 50)

   Display(root)

   printNthFromLast(root, 2)
}
