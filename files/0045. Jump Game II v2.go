package main

//leetcode 45. Jump Game II
//Greedy BFS version O(n)

import (
	"fmt"
	"time"
    "math/rand"
)
func jump(nums []int) int {
    steps,l,r := 0,0,0
    for r < len(nums)-1 {
        farthest := 0
        for i:=l;i<=r;i++ {
            if farthest < i+nums[i] {farthest = i+nums[i]}
        }
        l = r+1
        r = farthest
        steps += 1
    }
    return steps
}
func main() {
    nums := []int{2,3,1,1,4} //2
    fmt.Println(jump(nums))
    nums = []int{2,3,0,1,4} //2
    fmt.Println(jump(nums))
    n := 10_000
    nums = nil
    for i:=0;i<=n;i++ {
        nums = append(nums, rand.Intn(10))
    }
    start := time.Now()
    fmt.Println(jump(nums))
    end := time.Now()
    fmt.Println("Time:", end.Sub(start))

}
