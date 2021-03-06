package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))

	fmt.Println(combinationSum4([]int{1, 2, 3}, 32))
}

func combinationSum4(nums []int, target int) int {

	//dp[i][j] = count when j = nums[i] + num < nums[i]
	//dp[i][k] = dp[i][j] if k - j = nums[i]
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

func _combinationSum4(nums []int, target int) int {
	if len(nums) == 0 || target == 0 {
		return 0
	}
	sort.Ints(nums)

	df := make([][]int, len(nums))

	for i := range df {
		df[i] = make([]int, target+1)
		for j := range df[i] {
			df[i][j] = -1
		}
	}

	var process func(i int, target int) int

	process = func(i int, target int) int {
		if target < 0 {
			return 0
		}
		if df[i][target] >= 0 {
			return df[i][target]
		}
		if target == 0 {
			df[i][target] = 1
			return 1
		}

		r := 0
		for i, num := range nums {
			r += process(i, target-num)
		}
		df[i][target] = r
		return r
	}

	return process(0, target)
}
