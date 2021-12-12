package main
//leetcode 191. Number of 1 Bits

import (
    "fmt"
    "strconv"
)

func main () {
    n := "00000000000000000000000000001011"  //Output: 3
    // n := "00000000000000000000000010000000"  //Output: 1
    // n := "11111111111111111111111111111101"  //Output: 31
    binary, err := strconv.ParseUint(n, 2, 32)
    if err != nil {
        fmt.Println("Error during converting:", err)
    }
    bin32 := uint32(binary)
    fmt.Println(hammingWeight(bin32))

}

func hammingWeight(num uint32) int {
    var count int 
    for num > 0 {
        count += int(num&1)
        num >>= 1
    }
    return count
}
