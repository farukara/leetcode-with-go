package main
// leetcode 141. Linked List Cycle
// NOTE: results from this code does not correspond to the ones from leetcode because I did not make the cycled lists here. But if I did it would work fine. 
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
