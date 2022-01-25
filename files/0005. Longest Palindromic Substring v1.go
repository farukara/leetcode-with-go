package main
// 5. Longest Palindromic Substring 
// Approach-4 Expand around center approch - O(n2) time and O(1) space

import (
    "fmt"
)

func longestPalindrome(s string) string {
    subPalindrome := func (s string, i int, j int) string {
        counter := 0
        for i>=0 && j<len(s) && s[i] == s[j] {
            counter++
            i--
            j++
        }
        return s[i+1:j]
    }
    result := ""
    max := 0
    for i:=0;i<len(s);i++ {
        sub1 := subPalindrome(s,i,i)
        sub2 := subPalindrome(s,i,i+1)
        if len(sub1) > len(sub2) {
            if len(sub1) > max {
                max = len(sub1)
                result = sub1
            }
        } else if len(sub2) > max {
                max = len(sub2)
                result = sub2
        }
    }
    return result
}
func main() {
    s := "babad" //bab
    fmt.Println(longestPalindrome(s))
    s = "cbbd" //bb
    fmt.Println(longestPalindrome(s))
    s = "a" //a
    fmt.Println(longestPalindrome(s))
    s = "bb" //bb
    fmt.Println(longestPalindrome(s))
    s = "dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd" // time limit excedence
    fmt.Println(longestPalindrome(s))
}
