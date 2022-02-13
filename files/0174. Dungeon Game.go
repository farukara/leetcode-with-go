package main
//leetcode 174. Dungeon Game

import (
    "fmt"
    "math"
)
func calculateMinimumHP(dungeon [][]int) int {
    table := make([][]int, len(dungeon))
    mins := make([][]int, len(dungeon))
    table[0] = append(table[0], dungeon[0][0])
    mins[0] = append(mins[0], dungeon[0][0])
    for i:=1; i<len(dungeon); i++ {
        table[i] = append(table[i], table[i-1][0] + dungeon[i][0])
        if mins[i-1][0] < table[i][0] {
            mins[i] = append(mins[i], mins[i-1][0])
        } else {
            mins[i] = append(mins[i], table[i][0])
        }
    }

    for i:=1; i<len(dungeon[0]); i++ {
        table[0] = append(table[0], table[0][i-1] + dungeon[0][i])
        if mins[0][i-1] < table[0][i] {
            mins[0] = append(mins[0], mins[0][i-1])
        } else {
            mins[0] = append(mins[0], table[0][i])
        }
    }

    for i:=1; i<len(dungeon); i++ {
        for j:=1; j<len(dungeon[0]); j++ {
            var topless int
            var leftless int
            if mins[i-1][j] > 0 {
                topless = table[i-1][j]
            } else {
                if mins[i-1][j] < table[i-1][j] {
                    topless = int(math.Abs(float64(mins[i-1][j])))+1
                } else {
                    topless = table[i-1][j]
                }
            }
            if mins[i][j-1] > 0 {
                leftless = table[i][j-1]
            } else {
                if mins[i][j-1] < table[i][j-1] {
                    leftless = int(math.Abs(float64(mins[i][j-1])))+1
                } else {
                    leftless = table[i][j-1]
                }
            }
            // if math.Abs(float64(table[i-1][j])) < math.Abs(float64(table[i][j-1])) {
            if topless > leftless {
                table[i] = append(table[i], table[i-1][j] + dungeon[i][j])
                if mins[i-1][j] < table[i][j] {
                    mins[i] = append(mins[i], mins[i-1][j])
                } else {
                    mins[i] = append(mins[i], table[i][j])
                }
            } else {
                table[i] = append(table[i], table[i][j-1] + dungeon[i][j])
                if mins[i][j-1] < table[i][j] {
                    mins[i] = append(mins[i], mins[i][j-1])
                } else {
                    mins[i] = append(mins[i], table[i][j])
                }
            }
        }
    }

    fmt.Println(dungeon)
    fmt.Println(table)
    fmt.Println(mins)
    result := table[len(dungeon)-1][len(dungeon[0])-1]
    mask := result >>63
    result ^= mask
    result -=  mask
    result += 1
    min := mins[len(dungeon)-1][len(dungeon[0])-1]
    if min>0{return 1}
    return int(math.Abs(float64(min)))+1
}
func main() {
    dungeon := [][]int{{-2,-3,3},{-5,-10,1},{10,30,-5}} //7
    fmt.Println(calculateMinimumHP(dungeon))
    dungeon = [][]int{{0}} //1
    fmt.Println(calculateMinimumHP(dungeon))
    dungeon = [][]int{{100}} //1
    fmt.Println(calculateMinimumHP(dungeon))
    dungeon = [][]int{{0,5}, {-2,-3}} //1
    fmt.Println(calculateMinimumHP(dungeon))
}
