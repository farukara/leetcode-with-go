package main
//leetcode 0338. Counting Bits

import (
    "fmt"
)

func main() {
    fmt.Println(countBits(2))
    fmt.Println(countBits(5))
    fmt.Println(countBitsv2(2))
    fmt.Println(countBitsv2(5))
}

//naive O(n2)
func countBits(n int) []int {
    var ans []int
    total := 0
    for i:=0; i<=n; i++ {
        total = 0
        for current:=i;current>0; current>>=1 {
            total += current&1
        }
        ans = append(ans, total)
    }
    return ans
}

//dp version
func countBitsv2(n int) []int {
    ans := []int{0}
    // temp := 0
    for i:=1; i<=n; i++ {
        if i&1 == 1 {
            ans = append(ans, ans[i>>1]+1)
        } else {
            ans = append(ans, ans[i>>1])
        }
    }
    return ans
}
