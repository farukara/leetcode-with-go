package main
//leetcode 62. Unique Paths

import (
    "fmt"
)
func uniquePaths(m int, n int) int {
    table := make([][]int, m)
    for i:=0;i<m;i++ {
        table[i] = append(table[i], 1)
    }
    for i:=0;i<n;i++ {
        table[0] = append(table[0], 1)
    }
    for i:=1;i<m;i++ {
        for j:=1;j<n;j++ {
            table[i] = append(table[i], table[i-1][j] + table[i][j-1])
        }
    }
    return table[m-1][n-1]
}
func main() {
    m,n := 3,7 //28
    fmt.Println(uniquePaths(m,n))
    m,n = 3,2 //3
    fmt.Println(uniquePaths(m,n))
    
}
