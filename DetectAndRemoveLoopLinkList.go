package main

import "fmt"

type Node struct{
   next  *Node
   key   int
}

func NewNode(val int) *Node{
   return &Node{
      key:val,
      next: nil,
   }
}

func RemoveLoop(loop_node, head *Node){
   ptr1 := loop_node
   ptr2 := loop_node
   count := 1

   for ptr1.next != ptr2{
      count++
      ptr1 = ptr1.next
   }

   ptr1 = head
   ptr2 = head

   for i := 0; i < count; i++{
      ptr2 = ptr2.next
   }

   for ptr1 != ptr2{
      ptr1 = ptr1.next
      ptr2 = ptr2.next
   }

   for ptr2.next != ptr1 {
      ptr2 = ptr2.next
   }

   ptr2.next = nil
}

func DetectAndRemoveLoop(head *Node) bool{
   slow_p := head
   fast_p := head

   for (slow_p != nil) && (fast_p != nil) && (fast_p.next != nil){
      slow_p = slow_p.next
      fast_p = fast_p.next.next

      if slow_p == fast_p {
         RemoveLoop(slow_p, head)
         return true
      }
   }

   return false
}

func Display(root *Node) {
   for root != nil {
      fmt.Printf("%+v ->", root.key)
      root = root.next
   }

   fmt.Println()
}

func main(){
   head := NewNode(50)
   head.next = NewNode(20)
   head.next.next = NewNode(15)
   head.next.next.next = NewNode(4)
   head.next.next.next.next = NewNode(10)

   Display(head)

   fmt.Println(DetectAndRemoveLoop(head))
   head.next.next.next.next.next = head.next.next
   fmt.Println(DetectAndRemoveLoop(head))
   Display(head)
}
