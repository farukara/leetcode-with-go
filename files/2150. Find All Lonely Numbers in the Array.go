package main
//leetcode 2150. Find All Lonely Numbers in the Array

import (
    "fmt"
)

func findLonely(nums []int) []int {
    m := make(map[int]int)
    for _,v := range nums {
        if _,ok :=m[v]; ok {
            m[v]++
            fmt.Println(v, m[v])
        } else {
            m[v] = 1
        }
    }
    var lones []int
    for k := range m {
        _,l := m[k-1]
        _,h := m[k+1]
        if !l && !h && m[k] ==1 {
            lones = append(lones, k)
        }
    }
    return lones
}

func main () {
    nums := []int{10,6,5,8}  // {10,8}
    fmt.Println(findLonely(nums))

    nums = []int{1,3,5,3}  // k=5 {1,5}
    fmt.Println(findLonely(nums))
}
