package main

func singleNumber(nums []int) int {
	var res = 0

	for _, num := range nums {
		res ^= num
	}

	return res
}
