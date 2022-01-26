package main
//leetcode 53. Maximum Subarray

import (
    "fmt"
)

func maxSubArray(nums []int) int {
    sum := 0
    max := -1<<52
    for i:=0;i<len(nums);i++ {
        if nums[i] > sum + nums[i] {
            sum = nums[i]
            if sum > max {max=sum}
        } else {
            sum += nums[i]
            if sum > max {max=sum}
        }
    }
    return max
}

func main() {
    nums := []int{-2,1,-3,4,-1,2,1,-5,4} //6
    fmt.Println(maxSubArray(nums))
    nums = []int{1} //1
    fmt.Println(maxSubArray(nums))
    nums = []int{5,4,-1,7,8} //23
    fmt.Println(maxSubArray(nums))
    nums = []int{-1} //-1
    fmt.Println(maxSubArray(nums))
    nums = []int{-2, -1} //-1
    fmt.Println(maxSubArray(nums))
}
