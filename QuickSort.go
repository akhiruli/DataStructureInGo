package main

import "fmt"

func Swap(x *int, y *int) {
   tmp := *x
   *x = *y
   *y = tmp
}

func Partition(arr []int, low, high int) int{
   pivot := arr[high]
   i := low - 1
   for j := low; j < high; j++{
      if arr[j] < pivot{
         i++
         Swap(&arr[i], &arr[j])
      }
   }

   Swap(&arr[i + 1], &arr[high])
   return (i + 1)
}

func QuickSort(arr []int, low, high int){
   if low < high {
      pi := Partition(arr, low, high)
      QuickSort(arr, low, pi - 1)
      QuickSort(arr, pi + 1, high)
   }
}

func PrintArray(arr []int){
   for i:= 0; i < len(arr); i++{
      fmt.Printf("%d ", arr[i])
   }

   fmt.Println()
}

func main(){
   arr := []int{10, 7, 8, 9, 1, 5};
   QuickSort(arr, 0, len(arr) - 1)
   PrintArray(arr)
}
