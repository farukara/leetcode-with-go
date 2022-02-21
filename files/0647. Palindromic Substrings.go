package main
//leetcode 647. Palindromic Substrings - Expanding around center technique

import (
    "fmt"
)

func countSubstrings(s string) int {
    count := 0
    checkCenterAround := func(s string, i, j int) {
        for i>=0 && j<len(s) && s[i] == s[j] {
            count += 1
            i--
            j++
        }
    }
    for i:=0; i<len(s); i++ {
        checkCenterAround(s, i, i)
        checkCenterAround(s, i, i+1)
    }
    return count
}
func main() {
    s := "abc" //3
    fmt.Println(countSubstrings(s))
    s = "aaa" //6
    fmt.Println(countSubstrings(s))
}
