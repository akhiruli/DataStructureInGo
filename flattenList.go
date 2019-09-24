package main

import "fmt"

type Node struct{
   data  int
   next  *Node
   child *Node
}

func createList(arr []int, n int) *Node{
   var head, p *Node

   for i := 0; i < n; i++ {
      if head == nil{
         head = &Node{}
         p = head
      }else {
         p.next = &Node{};
         p = p.next;
      }
      p.data = arr[i];
      p.next = nil
      p.child = nil
   }
   return head;
}


func CreateList() *Node{
    arr1 := []int{10, 5, 12, 7, 11};
    arr2 := []int{4, 20, 13};
    arr3 := []int{17, 6};
    arr4 := []int{9, 8};
    arr5 := []int{19, 15};
    arr6 := []int{2};
    arr7 := []int{16};
    arr8 := []int{3};

    /* create 8 linked lists */
    head1 := createList(arr1, len(arr1));
    head2 := createList(arr2, len(arr2));
    head3 := createList(arr3, len(arr3));
    head4 := createList(arr4, len(arr4));
    head5 := createList(arr5, len(arr5));
    head6 := createList(arr6, len(arr6));
    head7 := createList(arr7, len(arr7));
    head8 := createList(arr8, len(arr8));


    /* modify child pointers to create the list shown above */
    head1.child = head2;
    head1.next.next.next.child = head3;
    head3.child = head4;
    head4.child = head5;
    head2.next.child = head6;
    head2.next.next.child = head7;
    head7.child = head8;


    /* Return head pointer of first linked list.  Note that all nodes are
       reachable from head1 */
    return head1;
}

func flattenList(head *Node) *Node{
   if head == nil {
      return head
   }

   tail := head
   for tail.next != nil{
      tail = tail.next
   }

   curr := head
   for curr != nil{
      if curr.child != nil{
         tail.next = curr.child
         temp := curr.child
         for temp.next != nil {
            temp = temp.next
         }

         tail = temp
      }

      curr.child = nil
      curr = curr.next
   }

   return head
}

func Display(root *Node) {
   for root != nil {
      fmt.Printf("%+v ->", root.data)
      root = root.next
   }

   fmt.Println()
}


func main(){
   head := CreateList()
   head = flattenList(head)

   Display(head)
}
