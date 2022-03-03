package main
// leetcode 206. Reverse Linked List

import (
    "fmt"
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

 func makeLList (a []int) *ListNode {
     if len(a) == 0 {return nil}
     head := &ListNode{Val: a[0], Next: nil}
     current := head
     for i:=1; i<len(a); i++ {
         node := &ListNode{Val: a[i]}
         current.Next = node
         current = node
     }
     return head
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {return nil}
    var fc *ListNode
    if head.Next == nil {
        return head
    } else {
        fc = head.Next
    }
    oldhead := head
    var sc *ListNode
    for head.Next != nil {
        tail := false
        if fc.Next !=nil {
            sc = fc.Next
        } else {
            tail =true
        }
        fc.Next = head
        head = fc
        if tail {
            break
        }
        fc = sc
    }
    oldhead.Next = nil
    return head
}
func main() {
    head := []int{1,2,3,4,5} // [5,4,3,2,1]
    fmt.Println("RESULT:", reverseList(makeLList(head)), "\n")
    head = []int{1,2} // [2,1]
    fmt.Println("result:", reverseList(makeLList(head)), "\n")
    head = []int{} // []
    fmt.Println("result:", reverseList(makeLList(head)), "\n")
}
