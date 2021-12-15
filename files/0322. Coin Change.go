package main
//leetcode 0322. Coin Change

import (
    "fmt"
)

func main() {
    coins := []int{1,2,5}
    amount := 11
    fmt.Println("expecting 3, got", coinChange(coins, amount))
    coins = []int{2}
    amount = 3
    fmt.Println("expecting -1, got", coinChange(coins, amount))
    coins = []int{1}
    amount = 0
    fmt.Println("expecting 0, got", coinChange(coins, amount))
    coins = []int{1}
    amount = 1
    fmt.Println("expecting 1, got", coinChange(coins, amount))
    coins = []int{1}
    amount = 2
    fmt.Println("expecting 2, got", coinChange(coins, amount))
    coins = []int{2,5,10,1}
    amount = 27
    fmt.Println("expecting 4, got", coinChange(coins, amount))
    coins = []int{186, 419, 83, 408}
    amount = 6249
    fmt.Println("expecting 20, got", coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
    if amount == 0 {return 0}
    table := make([][]int, len(coins)+1)
    for i:=0;i<=len(coins);i++ {
        table[i] = append(table[i], 0)
    }
    complement := 1^(1<<62)
    large_number := (1<<62)|complement
    for i:=1;i<=amount+1;i++ {
        table[0] = append(table[0], large_number)
    }

    for i:=1;i<len(coins)+1;i++ {
        for j:=1;j<amount+1;j++ {
            if j>=coins[i-1] {
                if table[i-1][j] < table[i][j-coins[i-1]]+1 {
                    table[i] = append(table[i], table[i-1][j])
                } else {
                    table[i] = append(table[i], table[i][j-coins[i-1]]+1)
                }
            } else {
                table[i] = append(table[i], table[i-1][j])
            }
        }
    }
    if table[len(coins)][amount] == 0 || table[len(coins)][amount]>large_number-1 {return -1}
    return table[len(coins)][amount]
}
