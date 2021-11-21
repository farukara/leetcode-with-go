package main
//leetcode 190. Reverse Bits

import (
    "fmt"
    "strconv"
)
func main() {
    // n := "00000010100101000001111010011100" //output: 964176192 
    n := "11111111111111111111111111111101"    //output: 3221225471
    binary, err := strconv.ParseUint(n, 2, 32)
    if err != nil {
        fmt.Println("Error during converting from binary:", err)
    }
    bin32 := uint32(binary)
    fmt.Println(reverseBits(bin32))
}

func reverseBits(num uint32) uint32 {
    var res uint32
    var digit uint32
    for i:=0; i<32; i++ {
        digit = num&1 
        num>>=1
        res<<=1 
        res |=digit
    }
    return res
}
