package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		var n int
		var S uint64
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &n)
		readUint64(s, pos+1, &S)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 3)
		}
		res := solve(n, int64(S), E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, S int64, E [][]int) int {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, 3)
	}

	for i := 0; i < len(E); i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	pq := make(PriorityQueue, 0, n)

	var tot int64

	var dfs func(p, u int) int

	dfs = func(p, u int) int {
		if len(adj[u]) == 1 && p >= 0 {
			return 1
		}

		var cnt int
		for _, i := range adj[u] {
			e := E[i]
			v, w := e[0]^e[1]^(u+1), e[2]
			v--
			if p == v {
				continue
			}
			c := dfs(u, v)
			cnt += c
			edge := NewEdge(w, c)

			tot += edge.priority()

			heap.Push(&pq, edge)
		}
		return cnt
	}

	dfs(-1, 0)

	var ans int
	for tot > S && pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Edge)
		tot -= cur.diff()
		cur.w /= 2
		ans++
		if cur.w > 0 {
			heap.Push(&pq, cur)
		}
	}
	return ans
}

type Edge struct {
	w     int
	c     int
	index int
}

func NewEdge(w, c int) *Edge {
	edge := new(Edge)
	edge.w = w
	edge.c = c
	return edge
}

func (edge Edge) priority() int64 {
	return int64(edge.w) * int64(edge.c)
}

func (edge Edge) diff() int64 {
	tot := edge.priority()
	return tot - int64(edge.w/2)*int64(edge.c)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].diff() > pq[j].diff()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Edge)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
