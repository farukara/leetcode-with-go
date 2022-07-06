package main
// leetcode 876. Middle of the Linked List

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

 func printLList (a *ListNode) {
    if a == nil {
        fmt.Print(" ")
        fmt.Println()
        return
    }
    fmt.Print(a.Val)
    fmt.Print(",")
    c := a
    for c.Next !=nil {
        c = c.Next
        fmt.Print(c.Val)
        fmt.Print(",")
    }
    fmt.Println()
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

func middleNode(head *ListNode) *ListNode {
    if head.Next == nil {
        return head
    }
    s:= head
    f:= head.Next
    for f.Next != nil && f.Next.Next != nil {
        f = f.Next.Next
        s = s.Next
    }
    return s.Next
}

func main() {
    a := []int{1,2,3,4,5} //{3,4,5}
    r := middleNode(makeLList(a))
    fmt.Println("input   : {1,2,3,4,5}")
    fmt.Println("expected: {3,4,5}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()

    a = []int{1,2,3,4,5,6} //{4,5,6}
    r = middleNode(makeLList(a))
    fmt.Println("input   : {1,2,3,4,5,6}")
    fmt.Println("expected: {4,5,6}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()

    a = []int{1,2} //{2}
    r = middleNode(makeLList(a))
    fmt.Println("input   : {1,2}")
    fmt.Println("expected: {2}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()
}
