/*Rearrange positive and negative numbers in O(n) time and O(1) extra space

An array contains both positive and negative numbers in random order. Rearrange the array elements so that positive and negative numbers are placed alternatively. Number of positive and negative numbers need not be equal. If there are more positive numbers they appear at the end of the array. If there are more negative numbers, they too appear in the end of the array.

For example, if the input array is [-1, 2, -3, 4, 5, 6, -7, 8, 9], then the output should be [9, -7, 8, -3, 5, -1, 2, 4, 6]

The solution is to first separate positive and negative numbers using partition process of QuickSort. In the partition process, consider 0 as value of pivot element so that all negative numbers are placed before positive numbers. Once negative and positive numbers are separated, we start from the first negative number and first positive number, and Swap every alternate negative number with next positive number.
*/
package main

import "fmt"

func Swap(x *int, y *int) {
   tmp := *x
   *x = *y
   *y = tmp
}
func Rearrange(arr []int){
   size := len(arr)
   i := -1
   /* The following few lines are similar to partition process
      of QuickSort.  The idea is to consider 0 as pivot and
      divide the array around it.*/

   for j := 0; j < size; j++{
      if arr[j] < 0 {
         i++
         Swap(&arr[i], &arr[j])
      }
   }

   pos := i + 1
   neg := 0

   for (pos < size) && (neg < pos) && arr[neg] < 0{
      Swap(&arr[neg], &arr[pos])
      neg = neg + 2
      pos++
   }
}

func PrintArray(arr []int){
   for i := 0; i < len(arr); i++{
      fmt.Printf("%d ", arr[i])
   }
   fmt.Println()
}

func main(){
   arr := []int{-1, 2, -3, 4, 5, 6, -7, 8, 9}
   Rearrange(arr)
   PrintArray(arr)
}
