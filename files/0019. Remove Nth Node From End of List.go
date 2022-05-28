package main
// leetcode 19. Remove Nth Node From End of List

import(
    "fmt"
    "sync"
)

 /**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
     Val int
     Next *ListNode
 }

 type LinkedList struct {
     head *ListNode
     size int
     lock sync.RWMutex
 }

 func makeLList (a []int) *ListNode {
     var result LinkedList
     var node ListNode
     if len(a) == 0 {
         return result.head
     }
     if len(a)== 1 {
         node.Val = a[0]
         node.Next =nil
         result.size = 1
         result.head = &node
         return result.head
     }
     node.Val = a[len(a)-1]
     node.Next = nil
     result.size = 1
     next := &node
     for i:=len(a)-2; i>=0; i-- {
         node := ListNode{a[i], next}
         next = &node
         result.size++
     }
     result.head = next
     return result.head
 }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    p := head
    c := head
    if n == 1 && head.Next == nil {
        head = nil
        return head
    }
    for i:=0; i<n-1 && c.Next != nil ; i++ {
        c = c.Next
    }
    if c.Next == nil {
        head = p.Next
        return head
    }
    c = c.Next
    for c.Next != nil {
        c = c.Next
        p = p.Next
    }
    p.Next = p.Next.Next
    return head
}

func main() {
    a := []int{1,2,3,4,5}
    n:=2  // {1,2,3,5}
    fmt.Println(removeNthFromEnd(makeLList(a), n))

    a = []int{1}
    n=1  // {}
    fmt.Println(removeNthFromEnd(makeLList(a), n))

    a = []int{1,2}
    n=1  // {1}
    fmt.Println(removeNthFromEnd(makeLList(a), n))
}
