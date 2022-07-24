package main
//leetcode 26. Remove Duplicates from Sorted Array

import (
    "fmt"
)

func removeDuplicates(nums []int) int {
    if len(nums) == 1 {return 1}
    s := 0;
    f := 1;
    for f<len(nums) {
        if nums[s] == nums[f] {
            f++
            if f>=len(nums) {break}
            for nums[s] == nums[f] {
                f++
                if f>=len(nums) {break}
            }
            if f>=len(nums) {break}
            s++
            nums[s] = nums[f]
            continue
        } else {
            s++
            f++
        }
    }
    return s+1
}

func main () {
    nums := []int{1,1,2}  // k=2 {1,2,_}
    fmt.Println(removeDuplicates(nums))

    nums = []int{0,0,1,1,1,2,2,3,3,4}  // k=5 {0,1,2,3,4,_,_,_,_,_}
    fmt.Println(removeDuplicates(nums))

    nums = []int{1,1,1}  // k=1 {1,_,_}
    fmt.Println(removeDuplicates(nums))
}
