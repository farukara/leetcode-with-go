package main
//leetcode 45. Jump Game II
//dp version O(n2)

import (
    "fmt"
    "time"
    "math/rand"
)
func jump(nums []int) int {
    var j []int
    j = append(j, 0)
    for i:=0;i<len(nums);i++ {
        for k:=1;k<=nums[i];k++ {
            if len(j)<=i+k {
                j = append(j, j[i]+1)
            } else {
                if j[i+k] > j[i] +1 {
                    j[i+k] = j[i] + 1
                }
            }
        }
    }
    return j[len(nums)-1]
}
func main() {
    nums := []int{2,3,1,1,4} //2
    fmt.Println(jump(nums))
    nums = []int{2,3,0,1,4} //2
    fmt.Println(jump(nums))
        n := 100_000
    nums = nil
    for i:=0;i<=n;i++ {
        nums = append(nums, rand.Intn(n))
    }
    start := time.Now()
    fmt.Println(jump(nums))
    end := time.Now()
    fmt.Println("Time:", end.Sub(start))

}
