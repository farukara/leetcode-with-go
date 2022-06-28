package main
//leetcode 27. Remove Element

import (
    "fmt"
)
func removeElement(nums []int, val int) int {
    if len(nums) == 0 {return 0}
    s,f := 0,0
    for f < len(nums) {
        if nums[f] != val {
            nums[s] = nums[f]
            s++
            f++
        } else {
            f++
        }
    }

    /* if nums[len(nums)-1] == val {
        if len(nums) == 1 {
            return 0
        }
        return s
    } else {
        if len(nums) == 1 {
            return 1
        }
        return s+1
    } */
    return s
}

func main () {
    a := []int{3,2,2,3}
    val := 3 // 2, {2,2, , }
    fmt.Println(removeElement(a,val), a)

    a = []int{0,1,2,2,3,0,4,2}
    val = 2 // 5, {0,1,3,0,4,_,_,_}
    fmt.Println(removeElement(a,val), a)

    a = []int{2}
    val = 3 // 1, {2}
    fmt.Println(removeElement(a,val), a)

    a = []int{3}
    val = 3 // 0, {3}
    fmt.Println(removeElement(a,val), a)

    a = []int{}
    val = 3 // 0, {}
    fmt.Println(removeElement(a,val), a)

    a = []int{3,3}
    val = 5 // 2, {3,3}
    fmt.Println(removeElement(a,val), a)
}
