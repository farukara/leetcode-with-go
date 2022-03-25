package main
//leetcode 442. Find All Duplicates in an Array

import (
    "fmt"
)

func findDuplicates(nums []int) []int {
    m := make(map[int]int)
    dups := make([]int, 0)
    for _,v := range nums {
        if _,ok := m[v]; ok {
            dups = append(dups, v)
        } else {
            m[v] = 1
        }
    }
    return dups
}

func main() {
    nums := []int{4,3,2,7,8,2,3,1} //[2,3]
    fmt.Println(findDuplicates(nums))
    nums = []int{1,1,2} // [1]
    fmt.Println(findDuplicates(nums))
    nums = []int{1} // []
    fmt.Println(findDuplicates(nums))
}
