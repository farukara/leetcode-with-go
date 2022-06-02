package main
// leetcode 24. Swap Nodes in Pairs

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

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    temp := head.Next
    p := head
    head.Next = head.Next.Next
    temp.Next = head
    head = temp
    c := head.Next.Next
    for c != nil {
        if c.Next == nil {
            return head
        }
        temp := c.Next
        p.Next = temp
        p = c
        c.Next = c.Next.Next
        temp.Next = c
        c = c.Next
    }
    return head
}

func main() {
    a := []int{1,2,3,4} //{2,1,4,3}
    r := swapPairs(makeLList(a))
    fmt.Println("input   : {1,2,3,4}")
    fmt.Println("expected: {2,1,4,3}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()

    a = []int{} //{}
    r = swapPairs(makeLList(a))
    fmt.Println("input   : {}")
    fmt.Println("expected: {}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()

    a = []int{1} // {1}
    r = swapPairs(makeLList(a))
    fmt.Println("input   : {1}")
    fmt.Println("expected: {1}")
    fmt.Print("got:       ")
    printLList(r)
    fmt.Println()
}
