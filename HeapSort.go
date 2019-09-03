package main

import "fmt"

func Swap(x *int, y *int) {
   tmp := *x
   *x = *y
   *y = tmp
}

func Heapify(A []int, n, keyIndex int){
   //Assuming array index starts from 0
   l := 2*keyIndex + 1 //left child
   r := 2*keyIndex + 2 //right child

   max := keyIndex

   if l < n && A[l] > A[max] {
      max = l
   }

   if r < n && A[r] > A[max] {
      max = r
   }

   if max != keyIndex{
      Swap(&A[keyIndex], &A[max])
      Heapify(A, n, max)
   }
}

func main(){
   A := []int{30, 15, 5, 20, 50}
   n := 5
   fmt.Println("Before Sort")
   for _, num := range A{
      fmt.Println(num)
   }
   for i := n/2 - 1; i >= 0; i--{
      Heapify(A, n, i)
   }

   for i := n - 1; i >= 0; i--{
      Swap(&A[0], &A[i])
      Heapify(A, i, 0)
   }

   fmt.Println("After Sort")
   for _, num := range A{
      fmt.Println(num)
   }
}
