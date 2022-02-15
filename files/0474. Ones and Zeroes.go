package main
//leetcode 474. Ones and Zeroes

import (
    "fmt"
)
func findMaxForm(strs []string, m int, n int) int {
    table := make([][]int, m+1)
    for i:=0;i<=m;i++ {
        for j:=0;j<=n;j++ {
            table[i] = append(table[i], 0)
        }
    }
    for _,s := range strs {
        one := 0
        zero := 0
        for _,l := range s {
            if string(l) == string("0") {
                zero += 1
            }
            if string(l) == string("1") {
                one += 1
            }
        }
        if zero<=m && one<=n {
            fmt.Println("s:", s)
            for i:=zero;i<=m;i++ {
            // for i:=m;i>m-zero;i-- {
                for j:=one;j<=n;j++ {
                // for j:=n;j>n-one;j-- {
                    fmt.Println("\n")
                    fmt.Println(table)
                    fmt.Println("i",i, "j", j, "zero", zero, "one", one)
                    fmt.Println("table i j:", table[i][j], "tablei-zero j-one:", table[i- zero][j- one] + 1) 
                    if table[i][j] < table[i- zero][j- one] + 1 {
                        table[i][j] = table[i- zero][j- one] + 1
                    }
                }
            }
        }
    }
    fmt.Println(table)
    return table[m][n]
    
}
func main() {
    m := 5
    n := 3
    strs := []string{"10","0001","111001","1","0"}
    fmt.Println(findMaxForm(strs, m, n))
}
