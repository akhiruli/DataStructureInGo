/*
count triplets with sum smaller than a given value

Given an array of distinct integers and a sum value. Find count of triplets with sum smaller than given sum value. Expected Time Complexity is O(n2).

Examples:

Input : arr[] = {-2, 0, 1, 3}
        sum = 2.
Output : 2
Explanation :  Below are triplets with sum less than 2
               (-2, 0, 1) and (-2, 0, 3) 
*/

package main

import(
   "fmt"
   "sort"
)

func CountTriplets(arr []int, sum int) int{
   // Sort input array
   sort.Ints(arr)

   n := len(arr)
   ans := 0
   for i := 0; i < n - 2; i++{
      j := i + 1
      k := n - 1
      for j < k{
         if (arr[i] + arr[j] + arr[k]) >= sum {
            k = k - 1
         }else{
            // This is important. For current i and j, there can be total k-j third elements
            ans += (k - j)
            j = j + 1
         }
      }
   }

   return ans
}

func main(){
   arr := []int{5, 1, 3, 4, 7}
   sum := 12;
   fmt.Println(CountTriplets(arr, sum))
}

