package main
//leetcode 237.Delete Node in a Linked List

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
func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}

func MakeLL (array []int) LList {
    var new_llist LList
    current := ListNode{array[0], nil}
    new_llist.head = &current
    for _, item := range array[1:] {
        next := ListNode{item, nil}
        if item == 4 {
            node = next
        }
        current.Next = &next
        current = next
    }
    return new_llist
}

type ListNode struct {
    Val int
    Next *ListNode
}

type LList struct {
    head *ListNode
}

var node ListNode

func main() {
    array := []int{1,2,3,4,5}
    myllist := MakeLL(array)
    fmt.Println(myllist)
}
