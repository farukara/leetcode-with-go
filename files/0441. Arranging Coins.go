package main
// leetcode 441. Arranging Coins

import "fmt"

func arrangeCoins(n int) int {
    s := 0
    for i:=1; i<=n; i++ {
        s++
        n -= i
        if n<0 {break}
    }
    return s
}

func main() {
    n := 5 //2
    fmt.Println(arrangeCoins(n))
    n = 8 //3
    fmt.Println(arrangeCoins(n))
}
