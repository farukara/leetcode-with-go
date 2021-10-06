package main
//leetcode 1. Two Sum

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    var result []int
    found := false
    for i:=0; i<len(nums)-1;i++ {
        for j:=i+1; j<len(nums); j++ {
            if nums[i] + nums[j] == target {
                result = append(result, i)
                result = append(result, j)
                found = true
            }
        }
        if found { break }
    }
    return result
}

func main() {
    data1 := []int {2,7,11,15}
    fmt.Println(twoSum(data1, 9))
    data2 := []int {3,2,4}
    fmt.Println(twoSum(data2, 6))
    data3 := []int {3,3}
    fmt.Println(twoSum(data3, 6))
}
