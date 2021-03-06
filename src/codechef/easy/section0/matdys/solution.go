package main

import "fmt"

var p2 []uint64

func init() {
	p2 = make([]uint64, 64)
	p2[0] = 1
	for i := 1; i < 64; i++ {
		p2[i] = p2[i-1] * 2
	}
}

func main() {
	var t int

	fmt.Scanf("%d", &t)

	for t > 0 {
		var n int
		var k uint64
		fmt.Scanf("%d %d", &n, &k)
		fmt.Println(solve(n, k))
		t--
	}
}

func solve(n int, k uint64) uint64 {
	arr := make([]uint64, n)

	for i := n - 1; i >= 0; i-- {
		arr[i] = k & 1
		k >>= 1
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	var res uint64

	for i := 0; i < n; i++ {
		res = res*2 + arr[i]
	}

	return res
}

func solve1(n int, k uint64) uint64 {
	k++
	var res uint64
	for i := n - 1; i >= 0; i-- {
		if k > p2[i] {
			res += p2[n-1-i]
			k -= p2[i]
		}
	}
	return res
}
