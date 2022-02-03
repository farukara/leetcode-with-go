package main
//leetcode 63. Unique Paths II

import (
    "fmt"
)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    table := make([][]int, m)
    if obstacleGrid[0][0] == 0 {
        table[0] = append(table[0], 1)
    } else {
        table[0] = append(table[0], 0)
    }
    for i:=1;i<m;i++ {
        if obstacleGrid[i][0] == 0 {
            table[i] = append(table[i], table[i-1][0])
        } else {
            table[i] = append(table[i], 0)
        }
    }
    for i:=1; i<n; i++ {
        if obstacleGrid[0][i] == 0 {
            table[0] = append(table[0], table[0][i-1])
        } else {
            table[0] = append(table[0], 0)
        }
    }
    for i:=1;i<m;i++ {
        for j:=1;j<n;j++ {
            if obstacleGrid[i][j] == 0 {
                table[i] = append(table[i], table[i-1][j] + table[i][j-1])
            } else {
                table[i] = append(table[i], 0)
            }
        }
    }
    return table[m-1][n-1]
}
func main() {
    obstacleGrid := [][]int{{0,0,0},{0,1,0},{0,0,0}} //2
    fmt.Println(uniquePathsWithObstacles(obstacleGrid))
    obstacleGrid = [][]int{{0,1},{0,0}} //1
    fmt.Println(uniquePathsWithObstacles(obstacleGrid))
    obstacleGrid = [][]int{{1}} //0
    fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
