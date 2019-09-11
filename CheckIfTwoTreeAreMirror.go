package main

import(
   "fmt"
)

type TreeNode struct {
   key      int
   left     *TreeNode
   right    *TreeNode
}

func NewNode(val int) *TreeNode{
   return &TreeNode{
      key: val,
      left: nil,
      right: nil,
   }
}

func IsMirror(t1, t2 *TreeNode) bool{
   if t1 == nil && t2 == nil {
      return true
   }

   if t1 == nil && t2 != nil {
      return false
   }

   if t2 == nil && t1 != nil {
      return false
   }

   return t1.key == t2.key && IsMirror(t1.left, t2.right) &&IsMirror(t1.right, t2.left)
}

func main(){
   a := NewNode(1)
   b := NewNode(1)
   a.left = NewNode(2)
   a.right = NewNode(3)
   a.left.left = NewNode(4)
   a.left.right = NewNode(5)

   b.left = NewNode(3)
   b.right = NewNode(2)
   b.right.left = NewNode(5)
   b.right.right = NewNode(4)

   ret := IsMirror(a, b)
   fmt.Printf("IsMirror:%v\n", ret)
}
