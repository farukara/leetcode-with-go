package main

import "fmt"

//leetcode 2. Add Two Numbers

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    c := 0
    var newnode ListNode
    t := &newnode
    ct := l1.Val + l2.Val
    t.Val = ct%10
    t.Next = nil
    c = ct /10
    current1 := l1
    current2 := l2
    currentt := t
    for current1.Next != nil && current2.Next != nil {
        current1 = current1.Next
        current2 = current2.Next
        ct = c + current1.Val + current2.Val
        var newnode ListNode
        currentt.Next = &newnode
        currentt = currentt.Next
        currentt.Val = ct%10
        currentt.Next = nil
        c = ct /10 
    }
    if current1.Next == nil && current2.Next == nil {
        for c>0 {
            var newnode ListNode
            currentt.Next = &newnode
            currentt = currentt.Next
            currentt.Val = c%10
            c /= 10
        }
        return t
    }
    var current *ListNode
    if current1.Next != nil {
        current = current1.Next
        var newnode ListNode
        currentt.Next = &newnode
        currentt = currentt.Next
        ct = c + current.Val
        currentt.Val = ct%10
        currentt.Next = nil
        c = ct /10
    } else if current2.Next != nil {
        current = current2.Next
        var newnode ListNode
        currentt.Next = &newnode
        currentt = currentt.Next
        ct = c + current.Val
        currentt.Val = ct%10
        currentt.Next = nil
        c = ct /10
    }
    for current.Next != nil {
        current = current.Next
        var newnode ListNode
        currentt.Next = &newnode
        currentt = currentt.Next
        ct = c + current.Val
        currentt.Val = ct%10
        currentt.Next = nil
        c = ct /10
    }
    for c>0 {
        var newnode ListNode
        currentt.Next = &newnode
        currentt = currentt.Next
        currentt.Val = c%10
        c /= 10
    }
    return t
}

func main() {
    l1 := []int{2,4,3}
    l2 := []int{5,6,4} //[7,0,8]
    res := addTwoNumbers(makeLList(l1), makeLList(l2))
    fmt.Print(res.Val)
    for res.Next != nil {
        res = res.Next
        fmt.Print(res.Val)
    }
    fmt.Println("--")
    l1 = []int{0}
    l2 = []int{0} //[0]
    res = addTwoNumbers(makeLList(l1), makeLList(l2))
    fmt.Print(res.Val)
    for res.Next != nil {
        res = res.Next
        fmt.Print(res.Val)
    }
    fmt.Println("--")
    l1 = []int{9,9,9,9,9,9,9}
    l2 = []int{9,9,9,9} // [8,9,9,9,0,0,0,1]
    res = addTwoNumbers(makeLList(l1), makeLList(l2))
    fmt.Print(res.Val)
    for res.Next != nil {
        res = res.Next
        fmt.Print(res.Val)
    }
    fmt.Println("--")
    l1 = []int{5}
    l2 = []int{5} // [0,5]
    res = addTwoNumbers(makeLList(l1), makeLList(l2))
    fmt.Print(res.Val)
    for res.Next != nil {
        res = res.Next
        fmt.Print(res.Val)
    }
    fmt.Println("--")
}
