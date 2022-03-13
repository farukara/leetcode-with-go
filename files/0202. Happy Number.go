package main
//leetcode 202. Happy Number

import (
    "fmt"
)
func isHappy(n int) bool {
    var m = make(map[int]struct{})
    for n !=1 {
        if n == 1 {return true}
        if _,ok := m[n]; ok {
            return false
        } else {
            m[n] = struct{}{}
        }
        var digits []int
        for n>9 {
            digits = append(digits, n%10)
            n = n/10
        }
        digits = append(digits, n%10)
        n = 0
        for _,v := range digits {
            n += v*v
        }
    }
    return true
}

func main() {
    n := 19 //true
    fmt.Println(isHappy(n))
    n = 2 //false
    fmt.Println(isHappy(n))
}
