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

func ReverseList(root *Node) *Node{
   if root == nil {
      return nil
   }

   var prev *Node
   curr := root
   var next *Node

   for curr != nil {
      next = curr.next
      curr.next = prev
      prev = curr
      curr = next
   }

   root = prev
   return root
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

   fmt.Println("Before Reverse:")
   Display(root)

   root = ReverseList(root)
   fmt.Println("After Reverse:")
   Display(root)
}
