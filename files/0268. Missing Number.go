package main
//letcode 0268. Missing Number

import (
    "fmt"
)

func main() {
    nums := []int{3,0,1}
    fmt.Println(missingNumber(nums))
    nums = []int{0,1}
    fmt.Println(missingNumber(nums))
    nums = []int{9,6,4,2,3,5,7,0,1}
    fmt.Println(missingNumber(nums))
    nums = []int{0}
    fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
    total := (len(nums)*(len(nums)+1))>>1
    for _,num := range nums {
        total -= num
    }
    return total
}
//if overflow is an issue than sum can be a running sum by: sum = sum+i+1-num
//another method is to use xor to cancel out each number with the counter, until the missing number remains
