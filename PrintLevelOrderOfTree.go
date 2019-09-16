package main

import "fmt"

type TreeNode struct{
   data     int
   left     *TreeNode
   right    *TreeNode
}

func NewNode(key int) *TreeNode{
   return &TreeNode{
      data: key,
      left: nil,
      right: nil,
   }
}

func Height(root *TreeNode) int{
   if root == nil{
      return 0
   }else{
      lh := Height(root.left)
      rh := Height(root.right)

      if lh > rh {
         return lh + 1
      }else{
         return rh + 1
      }
   }
}

func PrintAtGivenLevel(root *TreeNode, level int){
   if root == nil{
      return
   }else if level == 1{
      fmt.Printf("%d ", root.data)
   }else if level > 1{
      PrintAtGivenLevel(root.left, level - 1)
      PrintAtGivenLevel(root.right, level - 1)
   }
}

func PrintLevelOrder(root *TreeNode){
   h := Height(root)
   for i := 1; i<= h; i++{
      PrintAtGivenLevel(root, i)
      fmt.Println()
   }
}

func main(){
   root := NewNode(1)
   root.left = NewNode(2)
   root.right = NewNode(3)
   root.left.left = NewNode(4)
   root.left.right = NewNode(5)
   root.right.left = NewNode(6)
   root.right.right = NewNode(7)
   root.left.left.left = NewNode(8)
   root.left.left.right = NewNode(9)

   fmt.Printf("Height of Tree: %d\n",Height(root))
   fmt.Println("Level Order Traversal of the tree:")
   PrintLevelOrder(root)
}
