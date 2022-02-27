package main
//leetcode 448. Find All Numbers Disappeared in an Array

import (
    "fmt"
)

func findDisappearedNumbers(nums []int) []int {
    list := make(map[int]bool)
    for _,i := range nums {
        list[i] = true
    }
    var result []int
    for i:=1; i<=len(nums); i++ {
        _, ok := list[i]
        if !ok {
            result = append(result, i)
        }
    }
    return result
}
func main() {
    nums := []int{4,3,2,7,8,2,3,1} //5,6
    fmt.Println(findDisappearedNumbers(nums))
    nums = []int{1,1} //2
    fmt.Println(findDisappearedNumbers(nums))
}
