package main
//leetcode 258. Add Digits

import (
    "fmt"
)

func addDigits(num int) int {
    sum :=0
    for num>0 {
        sum += num%10
        num /= 10
    }
    if sum <= 9 {
        return sum
    } else {
        return addDigits(sum)
    }
}

func main() {
    num := 38 //2
    fmt.Println(addDigits(num))
    num = 0 //0
    fmt.Println(addDigits(num))
}
