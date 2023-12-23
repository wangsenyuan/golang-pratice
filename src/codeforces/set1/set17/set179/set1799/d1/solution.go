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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		cold := readNNums(reader, m)
		hot := readNNums(reader, m)
		res := solve(a, cold, hot)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a []int, cold []int, hot []int) int {
	n := len(a)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = 1e18
		}
	}

	dp[0][0] = 0
	dp[1][0] = cold[a[0]-1]

	for i := 2; i <= len(a); i++ {
		x := a[i-1]
		// dp[i-1][i] or dp[i][j]
		var y int = 1e18
		for j := 0; j < i-1; j++ {
			tmp := dp[i-1][j] + hot[x-1]
			if j == 0 || a[j-1] != x {
				tmp = dp[i-1][j] + cold[x-1]
			}
			y = min(y, tmp)
		}
		// dp[i][j] = dp[i][j-1] + cost
		for j := 0; j < i-1; j++ {
			tmp := dp[i-1][j] + hot[x-1]
			if a[i-2] != x {
				tmp = dp[i-1][j] + cold[x-1]
			}
			dp[i][j] = tmp
		}

		dp[i][i-1] = y
	}

	var ans int = 1e18

	for i := 0; i < n; i++ {
		ans = min(ans, dp[n][i])
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
