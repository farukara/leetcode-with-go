package main
//leetcode 7.Reverse Integer 

import (
    "fmt"
)

func reverse(x int32) int32 {
   if x < 0 {
       var byty []byte
       byty = byte(x)
       fmt.Println(byty)
   } 
   return byte
}

func main() {
    inputs := []int32{123, -123, 120}
    for _, input := range inputs {
        reverse(input)
    }
}
