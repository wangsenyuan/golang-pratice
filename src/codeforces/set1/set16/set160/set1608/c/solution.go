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
		n := readNum(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

func solve(A []int, B []int) string {
	n := len(A)
	// 如果x能赢y，y可以赢z，那么x就能赢z
	players := make([]Player, n)
	for i := 0; i < n; i++ {
		players[i] = Player{i, A[i], B[i]}
	}
	sort.Slice(players, func(i, j int) bool {
		return players[i].a < players[j].a
	})

	x := players[n-1].id

	g := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
	}

	for i := 0; i+1 < n; i++ {
		g[players[i].id][players[i+1].id] = true
	}
	sort.Slice(players, func(i, j int) bool {
		return players[i].b < players[j].b
	})

	// find the cycle with first player
	for i := 0; i+1 < n; i++ {
		g[players[i].id][players[i+1].id] = true
	}

	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = '0'
	}

	var dfs func(u int)

	dfs = func(u int) {
		ans[u] = '1'
		for v := range g[u] {
			if ans[v] == '0' {
				dfs(v)
			}
		}
	}

	dfs(x)
	dfs(players[n-1].id)

	return string(ans)
}

type Player struct {
	id int
	a  int
	b  int
}
