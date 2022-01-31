package main
//leetcode 55. Jump Game v2
//Greedy BFS O(n)

import (
    "fmt"
)

func canJump(nums []int) bool {
    l,r := 0,0 //left,right pointers
    for r<len(nums){
        k:=r
        for i:=l;i<=k;i++ {
            if nums[i]+i>r {r = i+nums[i]}
        }
        l=k+1
        if l>r {break}
    }
    if r>=len(nums)-1 {return true} else {return false}
}
func main() {
    nums := []int{2,3,1,1,4} //true
    fmt.Println(canJump(nums))
    nums = []int{3,2,1,0,4} //false
    fmt.Println(canJump(nums))
    nums = []int{0,2,3} //false
    fmt.Println(canJump(nums))
    nums = []int{3,0,8,2,0,0,1} //true
    fmt.Println(canJump(nums))
    nums = []int{2,0,2,0,1} //true
    fmt.Println(canJump(nums))
}
