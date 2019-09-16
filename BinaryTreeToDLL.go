/*
Given a Binary Tree (BT), convert it to a Doubly Linked List(DLL) In-Place. The left and right pointers in nodes are to be used as previous and next pointers respectively in converted DLL. The order of nodes in DLL must be same as Inorder of the given Binary Tree. The first node of Inorder traversal (left most node in BT) must be head node of the DLL.
*/

package main

import "fmt"

type Node struct{
   data     int
   left     *Node
   right    *Node
}

var prev *Node = nil

func BSTToDLL(root *Node, head **Node){
   if root == nil{
      return
   }

   BSTToDLL(root.left, head)

   if prev == nil{
      *head = root
   }else{
      root.left = prev
      prev.right = root
   }

   prev = root

   BSTToDLL(root.right, head)
}

func newNode(key int) *Node{
   return &Node{
      data: key,
      left: nil,
      right: nil,
   }
}

func printList(head *Node){
   for head != nil{
      fmt.Printf("%+v ->", head.data)
      head = head.right
   }
}

func main(){
   root := newNode(10);
   root.left        = newNode(12);
   root.right       = newNode(15);
   root.left.left  = newNode(25);
   root.left.right = newNode(30);
   root.right.left = newNode(36);

    // Convert to DLL
   var head *Node = nil
   BSTToDLL(root, &head);

    // Print the converted list
   printList(head);
   fmt.Println()
}
