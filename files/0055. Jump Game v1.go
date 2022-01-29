package main
//leetcode 55. Jump Game v1
//DP version O(n2)

import (
    "fmt"
)

func canJump(nums []int) bool {
    j := make([]int, len(nums)+1)
    j[0] = 1
    for i:=0;i<len(nums);i++ {
        if nums[i] == 0 {continue}
        if j[i] == 0 {continue}
        for k:=1;k<=nums[i];k++ {
            if i+k <len(nums) {
                j[i+k] = 1
            }
        }
    }
    if j[len(nums)-1] == 1 {
        return true
    }
    return false
}
func main() {
    nums := []int{2,3,1,1,4} //true
    fmt.Println(canJump(nums))
    nums = []int{3,2,1,0,4} //false
    fmt.Println(canJump(nums))
    nums = []int{0,2,3} //false
    fmt.Println(canJump(nums))
}
