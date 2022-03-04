package main
//leetcode 941. Valid Mountain Array

import (
    "fmt"
)
func validMountainArray(arr []int) bool {
    if len(arr) < 3 {return false}
    if arr[1] <=arr[0] {return false}
    inc := true
    index := 0
    for i:=1; i<len(arr); i++ {
        index = i
        if arr[i] > arr[i-1] {
            continue
        } else {
            inc = false
            break
        }
    }
    if inc {return false}
    for i:=index; i<len(arr); i++ {
        if arr[i] < arr[i-1] {
            continue
        } else {return false}
    }
    return true
}

func main() {
    arr := []int{2,1} //false
    fmt.Println(validMountainArray(arr))
    arr = []int{3,5,5} //false
    fmt.Println(validMountainArray(arr))
    arr = []int{0,3,2,1} //true
    fmt.Println(validMountainArray(arr))
    arr = []int{9,8,7,6,5,4,3,2,1,0} //false
    fmt.Println(validMountainArray(arr))
}
