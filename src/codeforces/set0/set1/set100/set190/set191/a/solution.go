package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = readString(reader)
	}
	res := solve(s)

	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(s []string) int {
	n := len(s)
	// dp[a][b] = 以a开始，以b结尾最长的距离
	dp := make([][]int, 26)
	for i := 0; i < 26; i++ {
		dp[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			dp[i][j] = -n * 10
		}
	}

	tmp := make([]int, 26)
	for i := 0; i < n; i++ {
		a := int(s[i][0] - 'a')
		b := int(s[i][len(s[i])-1] - 'a')
		for x := 0; x < 26; x++ {
			tmp[x] = dp[x][a] + len(s[i])
		}
		for x := 0; x < 26; x++ {
			dp[x][b] = max(dp[x][b], tmp[x])
		}
		dp[a][b] = max(dp[a][b], len(s[i]))
	}
	var res int

	for i := 0; i < 26; i++ {
		res = max(res, dp[i][i])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
