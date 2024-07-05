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
		p := readNNums(reader, n-1)
		res := solve(n, a, p)
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

type Pair struct {
	first  int
	second int
}

const inf = 1 << 30

func solve(n int, a []int, p []int) int {
	b := make([]int, n)

	for i := 0; i < n; i++ {
		b[i] = -a[i]
	}

	g := NewGraph(n, n)

	for i := 1; i < n; i++ {
		j := p[i-1] - 1
		b[j] += a[i]
		g.AddEdge(j, i)
	}

	merge := func(a []Pair, b []Pair) []Pair {
		// a and b both sorted
		if len(b) > len(a) {
			a, b = b, a
		}
		// len(a) >= len(b)
		for i := 0; i < len(b); i++ {
			x := b[len(b)-1-i]
			y := a[len(a)-1-i]
			a[len(a)-1-i] = Pair{x.first, x.second + y.second}
		}

		return a
	}

	var ans int

	var dfs func(u int, d int) []Pair

	dfs = func(u int, d int) []Pair {
		var res []Pair
		if g.nodes[u] == 0 {
			// a leaf
			res = append(res, Pair{d, inf})
			return res
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			tmp := dfs(v, d+1)
			res = merge(res, tmp)
		}

		if b[u] < 0 {
			// need to process this
			for i := len(res) - 1; i >= 0 && b[u] < 0; i-- {
				// 肯定可以处理完
				cur := res[i]
				x := min(-b[u], cur.second)
				ans += x * (cur.first - d)
				b[u] += x
				res[i] = Pair{cur.first, cur.second - x}
			}
		}
		// b[u] >= 0
		res = append(res, Pair{d, b[u]})
		return res
	}

	dfs(0, 0)

	return ans

}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
