package main
//leetcode 64. Minimum Path Sum

import (
    "fmt"
)
func minPathSum(grid [][]int) int {
    table := make([][]int, len(grid))
    table[0] = append(table[0], grid[0][0])
    for i:=1; i<len(grid); i++ {
        table[i] = append(table[i], table[i-1][0] + grid[i][0])
    }
    for i:=1; i<len(grid[0]); i++ {
        table[0] = append(table[0], table[0][i-1] + grid[0][i])
    }
    for i:=1; i<len(grid) ; i++ {
        for j:=1; j<len(grid[0]); j++ {
            if table[i-1][j] < table[i][j-1] {
                table[i] = append(table[i], table[i-1][j] + grid[i][j])
            } else {
                table[i] = append(table[i], table[i][j-1] + grid[i][j])
            }
        }
    }
    return table[len(grid)-1][len(grid[0])-1]
}
func main() {
    grid := [][]int{{1,3,1},{1,5,1},{4,2,1}} //7
    fmt.Println(minPathSum(grid))
    grid = [][]int{{1,2,3},{4,5,6}} //12
    fmt.Println(minPathSum(grid))
}
