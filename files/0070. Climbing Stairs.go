package main
//0070. Climbing Stairs

import (
    "fmt"
)
//
func climbStairs(n int) int {
    if n == 0 {return 0}
    if n == 1 {return 1}
    s2 := 1 // 2 previous stair
    s1 := 1 // 1 previous stair
    sc := 1 // current stair
    for i:=2;i<=n;i++ {
        sc = s1 + s2
        s2,s1 = s1,sc
    }
    return sc
}

func main() {
    fmt.Println(climbStairs(2)) //2
    fmt.Println(climbStairs(3)) //3
}
