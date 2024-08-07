package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	lessons := make([][]int, n)
	for i := 0; i < n; i++ {
		lessons[i] = readNNums(reader, 2)
	}

	res := solve(lessons)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

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

const MAX = 1_000_000 + 10

func solve(lessons [][]int) []int {
	tr := NewSegTree(MAX)
	for _, cur := range lessons {
		l, r := cur[0], cur[1]
		r--
		tr.Update(l, r, 1)
	}

	if tr.val[0] > 2 {
		return nil
	}
	n := len(lessons)

	if tr.val[0] == 1 {
		return all(n)
	}
	// tr.val[0] = 2
	var res []int
	for i, cur := range lessons {
		l, r := cur[0], cur[1]
		r--
		tr.Update(l, r, -1)
		if tr.val[0] == 1 {
			res = append(res, i+1)
		}
		tr.Update(l, r, 1)
	}

	return res
}

func all(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i + 1
	}
	return res
}

type SegTree struct {
	val  []int
	lazy []int
	sz   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &SegTree{arr, lazy, n}
}

func (tr *SegTree) pushValue(i int, v int) {
	tr.val[i] += v
	tr.lazy[i] += v
}

func (tr *SegTree) push(i int, l int, r int) {
	if l < r && tr.lazy[i] != 0 {
		tr.pushValue(2*i+1, tr.lazy[i])
		tr.pushValue(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		tr.push(i, l, r)
		if L <= l && r <= R {
			tr.pushValue(i, v)
			return
		}
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
		tr.val[i] = max(tr.val[2*i+1], tr.val[i*2+2])
	}
	loop(0, 0, tr.sz-1)
}