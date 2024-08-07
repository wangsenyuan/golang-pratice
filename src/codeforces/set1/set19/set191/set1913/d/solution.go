package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(a []int) int {
	n := len(a)

	dp := make([]int, n+1)
	pref := make([]int, n+1)
	stack := make([]int, n+1)
	var top int
	top++

	dp[0] = 1
	pref[0] = 1
	sum := 1
	for i := 1; i < n; i++ {
		for top > 0 && a[stack[top-1]] > a[i] {
			sum = sub(sum, dp[stack[top-1]])
			top--
		}
		if top == 0 {
			dp[i] = add(1, pref[i-1])
		} else {
			dp[i] = add(sum, sub(pref[i-1], pref[stack[top-1]]))
		}
		stack[top] = i
		top++
		sum = add(sum, dp[i])
		pref[i] = add(pref[i-1], dp[i])
	}

	var res int
	mn := 1 << 30
	for i := n - 1; i >= 0; i-- {
		mn = min(mn, a[i])
		if a[i] == mn {
			res = add(res, dp[i])
		}
	}

	return res
}

func solve1(a []int) int {
	n := len(a)
	sum := make([]int, n+2)
	dp := make([]int, n+2)

	sum[n] = 1
	stack := make([]int, n+2)
	var top int

	for i := n - 1; i >= 0; i-- {
		for top > 0 && a[stack[top-1]] > a[i] {
			top--
		}
		next := n
		if top > 0 {
			next = stack[top-1]
		}
		stack[top] = i
		top++
		cur := sub(sum[i+1], sum[next+1])
		if next != n {
			cur = add(cur, dp[next])
			dp[i] = add(sub(sum[next], sum[next+1]), dp[next])
		}
		sum[i] = add(cur, sum[i+1])
	}

	var res int
	mn := a[0]
	for i := 0; i < n; i++ {
		mn = min(mn, a[i])
		if a[i] == mn {
			res = add(res, sub(sum[i], sum[i+1]))
		}
	}

	return res
}
