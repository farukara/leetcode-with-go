package main
//leetcode 118. Pascal's Triangle

import (
    "fmt"
)
func generate(numRows int) [][]int {
    var r[][]int
    r = append(r, []int{1})
    for i:=2; i<=numRows; i++ {
        r = append(r, []int{})
        r[i-1] = append(r[i-1], 1)
        for j:=1;j<i-1;j++ {
            r[i-1] = append(r[i-1], r[i-2][j-1] + r[i-2][j])
        }
        r[i-1] = append(r[i-1], 1)
    }
    return r
}
func main() {
    numRows := 5
    fmt.Println(generate(numRows))
//Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

    numRows = 1
    fmt.Println(generate(numRows))
//Output: [[1]]
    
}
