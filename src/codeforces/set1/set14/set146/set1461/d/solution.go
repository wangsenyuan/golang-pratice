package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		Q := make([]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNum(reader)
		}
		res := solve(A, Q)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
	}
	fmt.Print(buf.String())
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(A []int, Q []int) []bool {
	sort.Ints(A)
	n := len(A)
	sum := make([]int64, n+1)
	for i, num := range A {
		sum[i+1] = sum[i] + int64(num)
	}

	cache := make(map[int64]bool)

	var dfs func(l int, r int)

	dfs = func(l int, r int) {
		if l == r {
			return
		}
		cache[sum[r]-sum[l]] = true
		tmp := (A[r-1] + A[l]) / 2
		mid := l
		for mid < r && A[mid] <= tmp {
			mid++
		}
		if mid < r {
			dfs(l, mid)
		}
		if l < mid {
			dfs(mid, r)
		}
	}

	dfs(0, n)

	ans := make([]bool, len(Q))

	for i, k := range Q {
		ans[i] = cache[int64(k)]
	}

	return ans
}
