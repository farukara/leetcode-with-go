package main
//leetcode 21. Merge Two Sorted Lists

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
        return nil
    }
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    result := list1
    if list1.Val < list2.Val {
        result = list1
        list1 = list1.Next
    } else {
        result = list2
        list2 = list2.Next
    }
    current := result
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            current = current.Next
            list1 = list1.Next
        } else {
            current.Next = list2
            current = list2
            list2 = list2.Next
        }
    }
    if list1 == nil {
        current.Next = list2
    }else {
        current.Next = list1
    }
    return result
}

func main() {
    list1 := []int{1,2,4}
    list2 := []int{1,3,4} //[1,1,2,3,4,4]
    result := mergeTwoLists(makeLList(list1), makeLList(list2))
    for result != nil {
        fmt.Println(result.Val)
        result = result.Next
    }
    list1 = []int{} 
    list2 = []int{} //[]
    fmt.Println(mergeTwoLists(makeLList(list1), makeLList(list2)))
    list1 = []int{}
    list2 = []int{0} //[0]
    fmt.Println(mergeTwoLists(makeLList(list1), makeLList(list2)))
}
