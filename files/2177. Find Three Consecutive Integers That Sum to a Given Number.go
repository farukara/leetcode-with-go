package main

import "fmt"

// leetcode 2177. Find Three Consecutive Integers That Sum to a Given Number

func sumOfThree(num int64) []int64 {
    if num%3 != 0 {return []int64{}}
    result := []int64{}
    result = append(result, num/3-1)
    result = append(result, num/3)
    result = append(result, num/3+1)

    return result
}

func main() {
    num := 33 //[10,11,12]
    fmt.Println(sumOfThree(int64(num)))
    num = 4 // []
    fmt.Println(sumOfThree(int64(num)))
}
