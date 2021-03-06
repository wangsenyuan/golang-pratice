package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	first := readNNums(scanner, 3)
	n, m, q := first[0], first[1], first[2]

	edges := make([][]int, m)

	for i := 0; i < m; i++ {
		edges[i] = readNNums(scanner, 3)
	}
	solver := NewSolver(n, edges)

	for q > 0 {
		q--

		scanner.Scan()
		bs := scanner.Bytes()

		var u, v int

		if bs[0] == '1' {
			pos := readInt(bs, 2, &u)
			readInt(bs, pos+1, &v)
			solver.AssigneZero(u, v)
		} else if bs[0] == '2' {
			pos := readInt(bs, 2, &u)
			readInt(bs, pos+1, &v)
			solver.AssignOriginal(u, v)
		} else {
			fmt.Println(solver.GetAnswer())
		}
	}
}

type Solver struct {
	n     int
	edges [][]int
	zeros *Node
	arr   []int
	cnt   []int
	mem   map[int]*Node
}

func find(arr []int, x int) int {
	if arr[x] != x {
		arr[x] = find(arr, arr[x])
	}
	return arr[x]
}

func union(arr []int, cnt []int, x, y int) bool {
	px := find(arr, x)
	py := find(arr, y)
	if px == py {
		return false
	}
	if cnt[px] >= cnt[py] {
		arr[py] = px
		cnt[px] += cnt[py]
	} else {
		arr[px] = py
		cnt[py] += cnt[px]
	}
	return true
}

func NewSolver(n int, edges [][]int) Solver {
	sort.Sort(Edges(edges))
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	used := make([][]int, 0, len(edges))

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		if union(arr, cnt, u, v) {
			used = append(used, edge)
		}
	}

	return Solver{n, used, new(Node), arr, cnt, make(map[int]*Node)}
}

func (solver *Solver) runMST() int {
	n := solver.n
	arr := solver.arr
	cnt := solver.cnt

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	cur := solver.zeros.next

	for cur != nil {
		u, v := cur.u, cur.v
		union(arr, cnt, u, v)
		cur = cur.next
	}

	if cnt[find(arr, 0)] == solver.n {
		return 0
	}

	var res int

	for _, edge := range solver.edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		if union(arr, cnt, u, v) {
			res += w
			if cnt[find(arr, u)] == solver.n {
				break
			}
		}
	}
	return res
}

type Node struct {
	u, v int
	next *Node
	prev *Node
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (solver *Solver) AssigneZero(u, v int) {
	u--
	v--

	key := solver.getKey(u, v)
	node := solver.mem[key]

	if node != nil {
		// already there
		return
	}

	node = new(Node)
	node.u = u
	node.v = v

	next := solver.zeros.next

	solver.zeros.next = node
	node.prev = solver.zeros
	if next != nil {
		node.next = next
		next.prev = node
	}

	solver.mem[key] = node
}

func (solver *Solver) getKey(u, v int) int {
	if u > v {
		u, v = v, u
	}
	return u*solver.n + v
}

func (solver *Solver) AssignOriginal(u, v int) {
	u--
	v--

	key := solver.getKey(u, v)

	node := solver.mem[key]

	if node == nil {
		return
	}
	prev, next := node.prev, node.next
	prev.next = next
	if next != nil {
		next.prev = prev
	}
	delete(solver.mem, key)
}

func (solver *Solver) GetAnswer() int {
	return solver.runMST()
}

type Edges [][]int

func (edges Edges) Len() int {
	return len(edges)
}

func (edges Edges) Less(i, j int) bool {
	a := edges[i]
	b := edges[j]
	return a[2] < b[2]
}

func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}
