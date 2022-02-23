package main
//leetcode 746. Min Cost Climbing Stairs

import (
    "fmt"
)
//
func minCostClimbingStairs(cost []int) int {
    c2 := 0 //cost at 2 ladders down
    c1 := cost[0] //cost at 1 ladder down
    cc := 0 //current cost
    for i:=2;i<len(cost)+1;i++ {
        if c1 + cost[i-1] < c2 + cost[i-1] {
            cc = c1 + cost[i-1]
        } else {
            cc = c2 + cost[i-1]
        }
        c2, c1 = c1, cc
    }
    if c1 < c2 {
        return c1
    } else {
        return c2
    }
}
func main() {
    cost := []int{10,15,20} //output:15
    fmt.Println(minCostClimbingStairs(cost))
    cost = []int{1,100,1,1,1,100,1,1,100,1} //output 6
    fmt.Println(minCostClimbingStairs(cost))
}
