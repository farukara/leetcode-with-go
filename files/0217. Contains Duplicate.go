package main

import "fmt"

func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})
    for _, v := range nums {
        _,ok := m[v]
        if ok {
            return true
        } else {
            m[v] = struct{}{}
        }
    }
    return false
}

func main() {
    nums := []int{1,2,3,1} //true
    fmt.Println(containsDuplicate(nums))
    nums = []int{1,2,3,4} //false
    fmt.Println(containsDuplicate(nums))
    nums = []int{1,1,1,3,3,4,3,2,4,2} //true
    fmt.Println(containsDuplicate(nums))
}
