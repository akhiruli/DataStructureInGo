package main

import "fmt"

func BinarySearch(A []int, key, l, h int) int{
   var mid int
   beg := l
   end := h
   for beg <= end {
      mid = (beg + end)/2
      if A[mid] == key {
         return mid
      } else if A[mid] < key{
         beg = mid + 1
      }else {
         end = mid - 1
      }
   }

   if A[mid] == key {
      return mid
   } else{
      return -1
   }
}

func main(){
   A := []int{10, 15, 25, 30, 50}

   index := BinarySearch(A, 25, 0, 4)
   fmt.Println(index)

   index = BinarySearch(A, 50, 0, 4)
   fmt.Println(index)
}
