package main
// leetcode 206. Reverse Linked List

import (
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
    result := reverseList(makeLList(head))
    for result != nil {
        fmt.Println(result.Val)
        result = result.Next
    }
    head = []int{1,2} // [2,1]
    result = reverseList(makeLList(head))
    fmt.Println("result:", reverseList(makeLList(head)), "\n")
    for result != nil {
        fmt.Println(result.Val)
        result = result.Next
    }
    head = []int{} // []
    result = reverseList(makeLList(head))
    fmt.Println("result:", reverseList(makeLList(head)), "\n")
    for result != nil {
        fmt.Println(result.Val)
        result = result.Next
    }
}
