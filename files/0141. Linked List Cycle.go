package main
// leetcode 141. Linked List Cycle
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
     var node ListNode
     if len(a) == 0 {
         // node.Val = nil
         node.Next = nil
         return &node
     }
     if len(a)== 1 {
         node.Val = a[0]
         node.Next =nil
         return &node
     }
     node.Val = a[len(a)-1]
     node.Next = nil
     prev := &node
     for i:=len(a)-2; i>=0; i-- {
         node.Val = a[i]
         prev.Next = &node
         prev = &node
     }
     return prev
 }
func hasCycle(head *ListNode) bool {
    if head == nil {return false}
    m := make(map[*ListNode]struct{})
    if head.Next ==nil {
        return false
    }
    m[head] = struct{}{}
    current := head
    for current.Next !=nil {
        current = current.Next
        if _,ok := m[current]; ok {
            return true
        } else {
            m[current] = struct{}{}
        }
    }
    return false
}

func main() {
    head := []int{3,2,0,-4}
    // pos := 1 //true
    fmt.Println(hasCycle(makeLList(head)))

    head = []int{1,2}
    // pos = 0 //true
    fmt.Println(hasCycle(makeLList(head)))

    head = []int{1}
    // pos = -1 //false
    fmt.Println(hasCycle(makeLList(head)))

    head = []int{}
    // pos = -1 //false
    fmt.Println(hasCycle(makeLList(head)))
}
