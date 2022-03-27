package main
//leetcode 1931. Painting a Grid With Three Different Colors

import (
    "fmt"
    // "math"
)
/* 1  2 1
1  3 1
2  1 2
2  3 2 
3  1 3
3  2 3

1  3 2
1  2 3
2  3 1
2  1 3
3  2 1
3  2 2 */

func colorTheGrid(m int, n int) int {
    table := make([][]int, m)
    table[0] = append(table[0], 3)
    for i:=1;i<m;i++ {
        table[i] = append(table[i], table[i-1][0]*2)
    }
    for i:=1;i<n;i++ {
        table[0] = append(table[0], table[0][i-1]*2)
    }
    
    for i:=1;i<m;i++ {
        for j:=1;j<n;j++ {
            // table[i] = append(table[i], (((table[i][j-1]*table[i-1][j])*2)- % int(math.Pow10(9)+7))
            table[i] = append(table[i], table[i-1][j]*3 + table[i][j-1]*3)
        }
    }
    fmt.Println(table)
    fmt.Println(m,n)
    return table[m-1][n-1]
}
func main() {
    m,n := 1, 1 //3
    fmt.Println(colorTheGrid(m,n))
    m,n = 1, 2 //6
    fmt.Println(colorTheGrid(m,n))
    m,n = 5, 5 //580986
    fmt.Println(colorTheGrid(m,n))
}
