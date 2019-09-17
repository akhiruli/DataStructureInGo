/*Leaders in an array

Write a program to print all the LEADERS in the array. An element is leader if it is greater than all the elements to its right side. And the rightmost element is always a leader. For example int the array {16, 17, 4, 3, 5, 2}, leaders are 17, 5 and 2.
*/

package main

import(
   "fmt"
)

/*inefficient method*/
func PrintLeaders(arr []int){
   size := len(arr)

   for i := 0; i < size ; i++{
      j := i + 1
      for ;j < size; j++{
         if arr[i] < arr[j]{
            break
         }
      }

      if j == size {
         fmt.Printf("%d ", arr[i])
      }
   }
}

/*Efficient Method*/
func printLeaders(arr []int){
   size := len(arr)
   max_from_right := arr[size - 1]
   fmt.Printf("%d ", max_from_right)
   for i := size - 2; i >= 0; i--{
      if max_from_right < arr[i]{
         max_from_right = arr[i]
         fmt.Printf("%d ", arr[i])
      }
   }
}

func main(){
   arr := []int{16, 17, 4, 3, 5, 2}
   PrintLeaders(arr)
   fmt.Println()
   printLeaders(arr)
   fmt.Println()
}
