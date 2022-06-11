package main
// leetcode 2095. Delete the Middle Node of a Linked List
// could be better implemented to avoid if branching by using f.next.next but this was first shot

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

func deleteMiddle(head *ListNode) *ListNode {
    s  := head //slow fast pointers
    f := head.Next
    if head.Next == nil {return nil}
    if head.Next.Next == nil {
    } else {
        var c int = 1
        f = f.Next
        for f != nil {
            printLList(head)
            fmt.Println("before c s f:", c, s.Val, f.Val)
            f = f.Next
            if c % 2 == 0 {
                s = s.Next
            }
            c++
            // fmt.Println("after c s f:", c, s.Val, f.Val)
        }
    }
    s.Next = s.Next.Next
    return head
}

func main() {
    a := []int{1,3,4,7,1,2,6} //{1,3,4,1,2,6}
    r := deleteMiddle(makeLList(a))
    fmt.Println("input   : {1,3,4,7,1,2,6}")
    fmt.Println("expected: {1,3,4,1,2,6}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()

    a = []int{1,2,3,4} //{1,2,4}
    r = deleteMiddle(makeLList(a))
    fmt.Println("input   : {1,2,3,4}")
    fmt.Println("expected: {1,2,4}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()

    a = []int{2,1} // {2}
    r = deleteMiddle(makeLList(a))
    fmt.Println("input   : {2,1}")
    fmt.Println("expected: {2}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()
}
