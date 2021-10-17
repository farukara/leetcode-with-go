package main
//leetcode 242.Valid Anagram

import (
    "fmt"
    "strings"
)

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false}
    var checked string
    for _, letter := range s {
        if strings.Contains(checked, string(letter)) {
            continue
        } else {
            if strings.Count(s, string(letter)) != strings.Count(t, string(letter)) {
                return false
            }
        }
        checked = checked + string(letter)
    }
    return true
    
}

func main() {
    fmt.Println(isAnagram("anagram", "nagaram"))
    fmt.Println(isAnagram("rat", "car"))

}
