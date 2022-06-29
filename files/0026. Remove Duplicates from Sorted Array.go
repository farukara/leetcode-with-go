package main
//leetcode 26. Remove Duplicates from Sorted Array

import (
    "fmt"
)

func removeDuplicates(nums []int) int {
    
}

func main () {
    nums := []int{1,1,2}  // k=2 {1,2,_}
    fmt.Println(removeDuplicates(nums))

    nums := []int{0,0,1,1,1,2,2,3,3,4}  // k=5 {0,1,2,3,4,_,_,_,_,_}
    fmt.Println(removeDuplicates(nums))
}
