package main

import (
	"fmt"
	"strconv"
)
func numDecodings(s string) int {
    if s[0] == '0' {return 0}
    count := 1
    for i:=1; i<len(s); i++ {
        if s[i] == '0' {
            count -=2
        } else if n,_ := strconv.Atoi(s[i-1:i+1]); n > 26 {
            fmt.Println("first")
            continue
        } else {
            fmt.Println("second")
            count += 1
        }
        if len(s)>=2 && s[1] == '0' {
            fmt.Println("final")
            count +=2
        }
        fmt.Println("\n")
    }
    return count
}
func main() {
    s := "12"
    fmt.Println(numDecodings(s)) //2
    s = "226"
    fmt.Println(numDecodings(s)) //3
    s = "06"
    fmt.Println(numDecodings(s)) //0
    s = "27"
    fmt.Println(numDecodings(s)) //1
    s = "2101"
    fmt.Println(numDecodings(s)) //1
    s = "10"
    fmt.Println(numDecodings(s)) //1
    s = "1123"
    fmt.Println(numDecodings(s)) //5
}
