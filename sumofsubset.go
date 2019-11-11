package main

import(
    "fmt"
)

func isSumOfSubset_dp(arr []int, n, sum int, dp []bool) bool{
   if sum == 0{
      return true
   }

   if n==0 && sum != 0{
      return false
   }

   if dp[n-1] == true{
      return true
   }

   ans := isSumOfSubset(arr, n-1, sum) || isSumOfSubset(arr, n-1, sum - arr[n-1])
   dp[n-1] = ans

   return ans
}

func isSumOfSubset(arr []int, n, sum int) bool{
    if sum == 0{
        return true
    }

    if n==0 && sum != 0{
        return false
    }


    return isSumOfSubset(arr, n-1, sum) || isSumOfSubset(arr, n-1, sum - arr[n-1])
}

func main(){
   arr := []int{3, 34, 4, 12, 5, 2}
   sum := 38

   fmt.Println(isSumOfSubset(arr, len(arr), sum))

   dp := make([]bool, len(arr))
   fmt.Println(isSumOfSubset_dp(arr, len(arr), sum, dp))
}
