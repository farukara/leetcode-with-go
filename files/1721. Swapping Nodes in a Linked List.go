
package main
// leetcode 1721. Swapping Nodes in a Linked List

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
    fmt.Print(a.Val)
    fmt.Print(",")
    c := a
    for c.Next !=nil {
        c = c.Next
        fmt.Print(c.Val)
        fmt.Print(",")
    }
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

func swapNodes(head *ListNode, k int) *ListNode {
    if head.Next == nil {return head}
    c:=1
    current := head
    one := head
    for current != nil {
        if c==k {
            one = current
        }
        current = current.Next
        c++
    }
    c--
    two := head
    limit := c-k+1
    /* if limit < 0 {
        limit = 0
    } */
    /* if k-1 > c/2 {
        limit = c-k+1
    } */
    for i:=1; i<limit; i++ {
        two = two.Next
    }
    one.Val, two.Val = two.Val, one.Val
    return head
}

func main() {
    a := []int{1,2,3,4,5}
    k:=2  // {1,4,3,2,5}
    r := swapNodes(makeLList(a), k)
    fmt.Println("input   : {1,2,3,4,5}")
    fmt.Println("expected: {1,4,3,2,5}")
    fmt.Print("           ")
    printLList(r)
    fmt.Println()

    a = []int{7,9,6,6,7,8,3,0,9,5}
    k=5  // {7,9,6,6,8,7,3,0,9,5}
    r = swapNodes(makeLList(a), k)
    fmt.Println("input   : {7,9,6,6,7,8,3,0,9,5}")
    fmt.Println("expected: {7,9,6,6,8,7,3,0,9,5}")
    fmt.Print("           ")
    printLList(r)
    fmt.Println()

    a = []int{100,90}
    k=2  // {90,100}
    r = swapNodes(makeLList(a), k)
    fmt.Println("input   : {100,90}")
    fmt.Println("expected: {90,100}")
    fmt.Print("           ")
    printLList(r)
    fmt.Println()

    a = []int{91,56,48,77,91,96,55,53,68}
    k=7  // {91,56,55,77,91,96,48,53,68}
    r = swapNodes(makeLList(a), k)
    fmt.Println("input   : {91,56,48,77,91,96,55,53,68}")
    fmt.Println("expected: {91,56,55,77,91,96,48,53,68}")
    fmt.Print("           ")
    printLList(r)
    fmt.Println()

    a = []int{1}
    k=1  // {91,56,55,77,91,96,48,53,68}
    r = swapNodes(makeLList(a), k)
    fmt.Println("input   : {1}")
    fmt.Println("expected: {1}")
    fmt.Print("           ")
    printLList(r)
    fmt.Println()

    a = []int{100,24,24,36,18,52,95,61,54,88,86,79,11,1,31,26}
    k=16  // {26,24,24,36,18,52,95,61,54,88,86,79,11,1,31,100}
    r = swapNodes(makeLList(a), k)
    fmt.Println("input   : {100,24,24,36,18,52,95,61,54,88,86,79,11,1,31,26}")
    fmt.Println("expected: {26,24,24,36,18,52,95,61,54,88,86,79,11,1,31,100}")
    fmt.Print("           ")
    printLList(r)
    fmt.Println()
}
