package main

import "fmt"

type TreeNode struct {
   key      int
   left     *TreeNode
   right    *TreeNode
}

func AddNodeToTree(root *TreeNode, val int) *TreeNode{
   if root == nil{
      return &TreeNode{
         key: val,
         left: nil,
         right: nil,
      }
   }

   if val < root.key {
      root.left = AddNodeToTree(root.left, val)
   }else if val > root.key{
      root.right = AddNodeToTree(root.right, val)
   }

   return root
}

func InOrderTraversal(root *TreeNode){
   if root != nil {
      InOrderTraversal(root.left)
      fmt.Printf("%+v ", root.key)
      InOrderTraversal(root.right)
   }
}

func main(){
   var root *TreeNode

   root = AddNodeToTree(root, 30)
   root = AddNodeToTree(root, 20)
   root = AddNodeToTree(root, 40)

   InOrderTraversal(root)
   fmt.Println()
}
