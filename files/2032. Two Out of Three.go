package main
//leetcode 2032. Two Out of Three

import (
    "fmt"
)

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    dic := make(map[int]int)
    internal_dic := make(map[int]int)
    all_set := [][]int{nums1, nums2, nums3}
    for _, numbs := range all_set {
        internal_dic = map[int]int{}
        if len(numbs) != 0 {
            for _, number := range numbs {
                _, err := internal_dic[number]
                if err {
                    continue
                } else {
                    internal_dic[number] = 1
                }
            }
        } else {
            continue
        }
        for k,_ := range internal_dic {
            _, err := internal_dic[k]
            if err {
                dic[k]++
            } else {
                dic[k] = 1
            }
        }
    }
    var result []int
    for k,v := range dic {
        if v >1 {
            result = append(result, k)
        }
    }
    return result
}

func main() {
    nums1 := []int{1,1,3,2}
    nums2 :=  []int{3,2}
    nums3 := []int{3}
    fmt.Println(twoOutOfThree(nums1, nums2, nums3))
    nums1 = []int{1,3}
    nums2 =  []int{3,2}
    nums3 = []int{1,2}
    fmt.Println(twoOutOfThree(nums1, nums2, nums3))
    nums1 = []int{1,2,2}
    nums2 =  []int{4,3,3}
    nums3 = []int{5}
    fmt.Println(twoOutOfThree(nums1, nums2, nums3))
}
