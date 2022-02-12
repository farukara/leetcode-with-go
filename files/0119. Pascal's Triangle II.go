
package main
//leetcode 119. Pascal's Triangle II

import (
    "fmt"
)
func getRow(rowIndex int) []int {
    var r[][]int
    r = append(r, []int{1})
    for i:=1; i<=rowIndex; i++ {
        r = append(r, []int{1})
        for j:=1;j<i;j++ {
            r[i] = append(r[i], r[i-1][j-1] + r[i-1][j])
        }
        r[i] = append(r[i], 1)
    }
    return r[rowIndex]
}
func main() {
    rowIndex := 3
    fmt.Println(getRow(rowIndex))
//Output: [1,3,3,1]

    rowIndex = 0
    fmt.Println(getRow(rowIndex))
//Output: [1]

    rowIndex = 1
    fmt.Println(getRow(rowIndex))
//Output: [1,1]
}
