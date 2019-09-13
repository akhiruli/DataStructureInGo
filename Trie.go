package main

import "fmt"

const A_SIZE=26

type TrieNode struct{
   children     [A_SIZE]*TrieNode; 
   isEndOfWord  bool
}

func GetNode() *TrieNode{
   node := new(TrieNode)

   for i := 0; i < A_SIZE; i++ {
      node.children[i] = nil
   }

   node.isEndOfWord = false

   return node
}

func Insert(root *TrieNode, key string) {
   crawl := root
   for i:= 0; i < len(key); i++ {
      index := key[i] - 'a'

      if crawl.children[index] == nil {
         crawl.children[index] = GetNode()
      }
      crawl = crawl.children[index]
   }

   crawl.isEndOfWord = true
}

func Search(root *TrieNode, key string) bool {
   crawl := root
   for i:= 0; i < len(key); i++ {
      index := key[i] - 'a'
      if crawl.children[index] == nil {
         return false
      }

      crawl = crawl.children[index]
   }

   return (crawl != nil && crawl.isEndOfWord)
}

func main(){
    keys := []string{"the", "a", "there",
                     "answer", "any", "by",
                     "bye", "their"}
   root := GetNode()

   for i := 0; i < len(keys); i++ {
      Insert(root, keys[i])
   }

   ret := Search(root, "the")
   fmt.Printf("Found: %v\n", ret)

   ret = Search(root, "answer")
   fmt.Printf("Found: %v\n", ret)

   ret = Search(root, "thetr")
   fmt.Printf("Found: %v\n", ret)
}
