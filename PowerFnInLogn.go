package main
import "fmt"

func Pow(x, y int) int{
   if y == 0{
      return 1
   }

   temp := Pow(x, y/2)
   if y % 2 == 0{
      return temp*temp
   }else{
      return x*temp*temp
   }
}

func PowIt(x, y int) int{
   res := 1
   for ;y > 0;{
      if y%2 != 0{
         res = res*x
      }
      y = y/2
      x = x*x
   }

   return res
}

func main(){
	fmt.Println(Pow(10,3))
	fmt.Println(PowIt(10,3))
}
