package main
//leetcode - 1646. Get Maximum in Generated Array

import (
    "fmt"
)
func getMaximumGenerated(n int) int {
    a := make([]int, n+1)
    if n == 0 {return 0}
    if n == 1 {return 1}
    max := 1
    a[0] = 0
    a[1] = 1
    for i:=1;i*2+1<=n;i++ {
        a[i*2] = a[i]
        if a[i*2] > max {max = a[i]}
        a[i*2+1] = a[i] + a[i+1]
        if a[i*2+1] > max {max = a[i*2+1]}
        fmt.Println(a)
    }
    return max
}
func main() {
    a:= []int{7,2,3}
    for _,input := range a {
        fmt.Println(getMaximumGenerated(input))
    }
}
