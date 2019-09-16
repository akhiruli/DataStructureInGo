/* Hashing implementation with collision resolution using Binary Search Tree
   Assuming all keys are unique*/
package main

import "fmt"
const HASH_TABLE_SIZE = 1000000
#const HASH_TABLE_SIZE = 10

type BSTNode struct{
   key      int
   data     interface{}
   left     *BSTNode
   right    *BSTNode
}

type Hash struct{
   hashTable [HASH_TABLE_SIZE]*BSTNode
}

func AddNodeToBST(root *BSTNode, key int, data interface{}) *BSTNode{
   if root == nil {
      return &BSTNode{
         key: key,
         data: data,
         left: nil,
         right: nil,
      }
   }

   if key < root.key {
      root.left = AddNodeToBST(root.left, key, data)
   }else{
      root.right = AddNodeToBST(root.right, key, data)
   }

   return root
}

func FindInBST(root *BSTNode, key int) interface{}{
   if root == nil {
      return nil
   }

   if root.key == key{
      return root.data
   }

   if key < root.key {
      return FindInBST(root.left, key)
   }

   return FindInBST(root.right, key)
}

func (h *Hash) Add(key int, data interface{}) {
   index := key % HASH_TABLE_SIZE
   h.hashTable[index] = AddNodeToBST(h.hashTable[index], key, data)
}

func (h *Hash) Find(key int) interface{}{
   index := key % HASH_TABLE_SIZE
   return FindInBST(h.hashTable[index], key)
}

func main(){
   hash := new(Hash)
   v1 := 200000
   v2 := 300000
   v3 := 400000
   v4 := 1100000

   hash.Add(1, v1)
   hash.Add(2, v2)
   hash.Add(3, v3)
   hash.Add(11, v4)

   fmt.Printf("first: %v\n", hash.Find(11))
}
