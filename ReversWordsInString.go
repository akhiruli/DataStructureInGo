package main

import(
   "fmt"
   "strings"
)

func reverse_words(str string) string{
   words := strings.Fields(str)
   for i, j := 0, len(words) - 1; i < j; i, j = i + 1, j - 1{
      words[i], words[j] = words[j], words[i]
   }

   return strings.Join(words, " ")
}

func main(){
   str := "I will go to school"
   fmt.Println(str)
   fmt.Println(reverse_words(str))
}
