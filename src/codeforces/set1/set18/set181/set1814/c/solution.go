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
		n, x, y := readThreeNums(reader)
		r := readNNums(reader, n)
		a, b := solve(x, y, r)

		buf.WriteString(fmt.Sprintf("%d ", len(a)))
		for i := 0; i < len(a); i++ {
			buf.WriteString(fmt.Sprintf("%d ", a[i]))
		}
		buf.WriteString(fmt.Sprintf("\n%d ", len(b)))
		for i := 0; i < len(b); i++ {
			buf.WriteString(fmt.Sprintf("%d ", b[i]))
		}
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

func solve(s1 int, s2 int, r []int) ([]int, []int) {
	n := len(r)
	tasks := make([]Pair, n)

	for i := 0; i < n; i++ {
		tasks[i] = Pair{r[i], i}
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].first > tasks[j].first
	})

	var a []int
	var b []int

	r1, r2 := Pair{0, 0}, Pair{0, 1}

	update := func(next int, id int) {
		if r1.second == id {
			r1 = Pair{next, id}
		} else {
			r2 = Pair{next, id}
		}
		if r1.first > r2.first {
			r1, r2 = r2, r1
		}
	}

	update(s1, 0)
	update(s2, 1)

	for _, cur := range tasks {
		if r1.second == 0 {
			a = append(a, cur.second+1)
			next := r1.first + s1
			update(next, 0)
		} else {
			b = append(b, cur.second+1)
			next := r1.first + s2
			update(next, 1)
		}
	}

	return a, b
}

type Pair struct {
	first  int
	second int
}
