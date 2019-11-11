/*
   In a given binary tree, we need to find maximum score starting from root to leave
   score calculation is summation of all the value of the node visited.
   eg.
            2
          /   \
          5   15
         /
         5
        /
       15

   Answer should be : 2 + 5 + 5 + 15 = 27      
*/
package main

import (
   "fmt"
)

type Node struct{
   data     int
   left     *Node
   right    *Node
}

func newNode(k int) *Node{
   return &Node{
      data: k,
      left: nil,
      right: nil,
   }
}

func GetTotalNodes(root *Node) int{
   if root == nil{
      return 0
   }

   return 1 + GetTotalNodes(root.left) + GetTotalNodes(root.right)
}

func max(n1, n2 int) int{
   if n1 >= n2{
      return n1
   }else{
      return n2
   }
}

func GetMaxScore(root *Node) int{
   if root == nil{
      return 0
   }

   if root.left == nil && root.right == nil{
      return root.data
   }

   return root.data + max(GetMaxScore(root.left), GetMaxScore(root.right))
}

func GetMaxScore_dp(root *Node, dp []int, n int) int{
   if root == nil{
      return 0
   }

   if root.left == nil && root.right == nil{
      return root.data
   }

   if dp[n] != 0{
      return dp[n]
   }

   dp[n] = root.data + max(GetMaxScore_dp(root.left, dp, n+1), GetMaxScore_dp(root.right, dp, n+1))

   return dp[n]

}

func main(){
   root := newNode(2)
   root.left = newNode(5)
   root.right = newNode(15)
   root.left.left = newNode(5)
   root.left.left.left = newNode(15)

   fmt.Println(GetMaxScore(root))
   total_nodes := GetTotalNodes(root)

   //fmt.Println(total_nodes)

   dp := make([]int, total_nodes)

   fmt.Println(GetMaxScore_dp(root, dp, 0))
}
