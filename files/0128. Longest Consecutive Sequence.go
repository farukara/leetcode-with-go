package main

import "fmt"

//leetcode 128. Longest Consecutive Sequence

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {return 0}
    m := make(map[int]struct{})
    var max int
    var maxinner int
    for _,n := range nums {
        m[n] = struct{}{}
    }
    for n := range m {
        if _,ok := m[n-1]; ok {
            continue
        }
        ok := true
        maxinner = 0
        current := n
        for ok {
            maxinner += 1
            current += 1
            _,ok = m[current]
        }
        if maxinner>max {max=maxinner}
    }
    return max
}

func main() {
    nums := []int{100,4,200,1,3,2} //4
    fmt.Println(longestConsecutive(nums))
    nums = []int{0,3,7,2,5,8,4,6,0,1} //9
    fmt.Println(longestConsecutive(nums))
    nums = []int{9,1,4,7,3,-1,0,5,8,-1,6} //7
    fmt.Println(longestConsecutive(nums))
}
