package main

import (
   "fmt"
   "sync"
)

func MakeInfinite()(chan<-interface{}, <-chan interface{}){
   in := make(chan interface{})
   out := make(chan interface{})

   go func(){
      var inQueue []interface{}
      outCh := func() chan interface{}{
         if len(inQueue) == 0 {
            return nil
         }
         return out
      }

      churVal := func() interface{}{
         if len(inQueue) == 0 {
            return nil
         }

         return inQueue[0]
      }

      for len(inQueue) > 0 || in != nil{
         select{
            case v, ok := <-in:
               if !ok{
                  in = nil
               }else{
                  inQueue = append(inQueue, v)
               }
            case outCh() <- churVal():
               inQueue = inQueue[1:]
         }
      }
      close(out)
   }()
   return in, out
}

func main(){
   in, out := MakeInfinite()
   lastVal := -1
   var wg sync.WaitGroup
   wg.Add(1)
   go func(){
      for v := range out{
         vi := v.(int)
         fmt.Println(vi)
         if lastVal + 1 != vi {
            fmt.Printf("Unexpected value: expected %d, got %d\n", lastVal + 1, vi)
         }

         lastVal = vi
      }
      wg.Done()
      fmt.Println("finished reading")
   }()

   for i:=0; i<100;i++{
      fmt.Println("writing", i)
      in <-i
   }

   close(in)
   fmt.Println("Finished writing")
   wg.Wait()

   if lastVal != 99 {
      fmt.Printf("Didn't get all values, last one received was %d\n", lastVal)
   }
}
