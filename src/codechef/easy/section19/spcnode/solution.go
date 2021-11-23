package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(y int) int {
		fmt.Printf("? %d\n", y)
		return readNum(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		x := solve(n, E, ask)

		fmt.Printf("! %d\n", x)
	}

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
func solve(n int, E [][]int, ask func(int) int) int {
	g := NewGraph(n, len(E)*2)
	for i := 0; i < len(E); i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	marked := make([]bool, n)
	parent := make([]int, n)
	size := make([]int, n)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		parent[u] = p
		size[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || marked[v] {
				continue
			}
			dfs(u, v)
			size[u] += size[v]
		}
	}

	var centroid func(p, u int, tot int) int

	centroid = func(p, u int, tot int) int {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || marked[v] {
				continue
			}
			if size[v]*2 > tot {
				return centroid(u, v, tot)
			}
		}
		return u
	}

	var start int
	for {
		dfs(-1, start)
		c := centroid(-1, start, size[start])
		q := ask(c + 1)
		if q < 0 {
			return c + 1
		}
		q--
		marked[c] = true
		dfs(-1, c)

		for parent[q] != c {
			q = parent[q]
		}
		start = q
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
