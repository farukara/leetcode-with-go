package main
// leetcode 371. Sum of Two Integers

import "fmt"

func getSum(a int, b int) int {
    t := a^b
    c := (a&b) << 1
    for c !=0 {
        temp := t
        t ^= c
        c = (c&temp) <<1
    }
    return t
}

func main() {
    a,b := 1,2 //3
    fmt.Println(getSum(a,b))
    a,b = 2,3 //5
    fmt.Println(getSum(a,b))
    a,b = -112,-5 //-117
    fmt.Println(getSum(a,b))
    a,b = 12,-112 //-100
    fmt.Println(getSum(a,b))
    a,b = -2,3 //1
    fmt.Println(getSum(a,b))
}
