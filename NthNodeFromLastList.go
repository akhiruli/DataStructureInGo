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
   Maintain two pointers – reference pointer and main pointer. Initialize both reference and main pointers to head. First move reference pointer to n nodes from head. Now move both pointers one by one until reference pointer reaches end. Now main pointer will point to nth node from the end. Return main pointer.
*/

func GetNthNodeFromLast(root *Node, n int) *Node{
   main_ptr := root
   ref_ptr := root

   count := 0
   for count < n {
      if ref_ptr == nil {
         return nil
      }

      ref_ptr = ref_ptr.next
      count++
   }

   for ref_ptr != nil {
      ref_ptr = ref_ptr.next
      main_ptr = main_ptr.next
   }

   return main_ptr
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

   temp := GetNthNodeFromLast(root, 2)
   fmt.Println(temp.key)
}
