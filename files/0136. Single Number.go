package main

import (
    "fmt"
)

func main() {
    nums := []int{2,2,1}
    fmt.Println(singleNumber(nums))
    nums = []int{4,1,2,1,2}
    fmt.Println(singleNumber(nums))
    nums = []int{1}
    fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
    target := 0
    for _,num := range nums {
        target ^= num
    }
    return target
}
