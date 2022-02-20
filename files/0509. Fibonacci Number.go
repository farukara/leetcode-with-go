package main

import (
    "fmt"
)

func fib(n int) int {
    if n == 0 {return 0}
    if n == 1 {return 1}
    fib1 := 1 //previous
    fib2 := 0 //2 previous
    fibc := 0 //current
    for i:=2;i<=n;i++ {
        fibc = fib2+fib1
        fib2, fib1 = fib1, fibc
    }
    return fibc
}
func main() {
    a := []int{2,3,4}
    for _,input := range a {
        fmt.Println(fib(input))
    }
    
}
