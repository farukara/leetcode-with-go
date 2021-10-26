package main
//leetcode 7.Reverse Integer 

import (
    "fmt"
    "strconv"
    "math"
)

func reverse(x int32) int32 {
    var remainder int32
    var result string
   if x > 0 {
       for {
           if x>9 {
               remainder = x%10
               x = math.Floor(x/10)
           } else {
               result = append(result)

           }
       }
   } 
   return 32
}

func main() {
    inputs := []int32{123, -123, 120, 121, 122}
    for _, input := range inputs {
        reverse(input)
    }
}
