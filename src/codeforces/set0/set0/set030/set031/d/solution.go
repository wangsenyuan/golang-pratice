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
	W, H, n := readThreeNums(reader)
	cuts := make([][]int, n)
	for i := 0; i < n; i++ {
		cuts[i] = readNNums(reader, 4)
	}
	res := solve(W, H, cuts)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

func solve(W int, H int, breaks [][]int) []int {
	// can[.][.][0] 表示可以在水平方向能够移动
	// can[.][.][1]表示在垂直方向可以移动
	var hori, vert [100][100]bool

	for _, cur := range breaks {
		if cur[0] == cur[2] {
			for y := cur[1]; y < cur[3]; y++ {
				// 它不能
				vert[cur[0]][y] = true
			}
		} else {
			for x := cur[0]; x < cur[2]; x++ {
				hori[x][cur[1]] = true
			}
		}
	}

	var mark [100][100]bool

	var dfs func(w int, h int) int

	dfs = func(w int, h int) int {
		mark[w][h] = true
		res := 1
		if h+1 < H && !mark[w][h+1] && !hori[w][h+1] {
			res += dfs(w, h+1)
		}
		if h-1 >= 0 && !mark[w][h-1] && !hori[w][h] {
			res += dfs(w, h-1)
		}
		if w+1 < W && !mark[w+1][h] && !vert[w+1][h] {
			res += dfs(w+1, h)
		}
		if w-1 >= 0 && !mark[w-1][h] && !vert[w][h] {
			res += dfs(w-1, h)
		}
		return res
	}
	var res []int
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			if !mark[i][j] {
				res = append(res, dfs(i, j))
			}
		}
	}
	sort.Ints(res)

	return res
}
