package main
//leetcode 202. Happy Number

import (
    "fmt"
)
func isHappy(n int) bool {
    ss := func (i int) int {
        tot := 0
        for i > 0 {
            tot += (i%10) * (i%10)
            i /= 10
        }
        return tot
    }
    slow := ss(n)
    fast := ss(ss(n))
    for slow != fast {
        if slow == 1 || fast == 1 {return true}
        slow = ss(slow)
        fast = ss(ss(fast))
    }
    return false
}

func main() {
    n := 19 //true
    fmt.Println(isHappy(n))
    n = 2 //false
    fmt.Println(isHappy(n))
}
