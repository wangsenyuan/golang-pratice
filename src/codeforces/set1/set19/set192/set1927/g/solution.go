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

func solve(a []int) int {
	n := len(a)
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, n+1)
			for k := 0; k <= n; k++ {
				dp[i][j][k] = n
			}
		}
	}
	dp[0][0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k <= n; k++ {
				// 如果不使用i处的charge
				// 移动到i+1时
				// 正常nj要增加1，但是如果在左边没有被覆盖到，那么还是要没有覆盖
				// 但是如果右边原来没有覆盖到(k=0), 那么此时左边被覆盖的cells的个数，是1（因为i肯定是被覆盖到的）
				// 向右边移动时，k要递减
				ni := i + 1
				var nj int
				if j > 0 {
					nj = j + 1
				} else if k == 0 {
					nj = 1
				}
				nk := max(k-1, 0)
				dp[ni][nj][nk] = min(dp[ni][nj][nk], dp[i][j][k])

				// 使用i向右charge
				nj = 0
				if j > 0 {
					nj = j + 1
				}
				// 那么在i的范围内，都被覆盖到了
				if nj <= a[i] {
					nj = 0
				}
				dp[ni][nj][nk] = min(dp[ni][nj][nk], dp[i][j][k]+1)

				nj = 0
				if j > 0 {
					nj = j + 1
				}

				nk = max(a[i]-1, k-1)
				dp[ni][nj][nk] = min(dp[ni][nj][nk], dp[i][j][k]+1)
			}
		}
	}
	res := n
	for k := 0; k <= n; k++ {
		res = min(res, dp[n][0][k])
	}
	return res
}
