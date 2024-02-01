package main

import (
   "fmt"
)


func divideArray(nums []int, k int) [][]int {
    sort.Ints(nums)
    if len(nums) % 3 != 0 {
        return nil
    }
    var res [][]int
    for i := 0; i < len(nums); i += 3 {
        if nums[i+2] - nums[i] > k {
            return nil
        }
        res = append(res, nums[i:i+3])
    }
    return res
}
