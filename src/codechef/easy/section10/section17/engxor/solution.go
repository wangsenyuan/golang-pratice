package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func main() {
	scanner := bufio.NewReader(os.Stdin)

	tc := readNum(scanner)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, q := readTwoNums(scanner)

		A := readNNums(scanner, n)

		solver := NewSolver(n, A)

		for q > 0 {
			q--
			p := readNum(scanner)
			a, b := solver.Ask(p)
			buf.WriteString(fmt.Sprintf("%d %d\n", a, b))
		}
	}
	fmt.Print(buf.String())
}

const MAX_N = 100001

type Solver struct {
	even, odd int
}

func NewSolver(n int, A []int) Solver {
	var odd, even int

	for i := 0; i < n; i++ {
		var cnt int
		for A[i] > 0 {
			cnt += A[i] & 1
			A[i] >>= 1
		}
		if cnt&1 == 1 {
			odd++
		} else {
			even++
		}
	}

	return Solver{odd, even}
}

func (solver Solver) Ask(p int) (x int, y int) {
	var cnt int
	for p > 0 {
		cnt += p & 1
		p >>= 1
	}
	// o + o => e
	// o + e => o
	// e + e => e
	// e + o => o
	if cnt&1 == 1 {
		// odd
		x = solver.even
		y = solver.odd
	} else {
		x = solver.odd
		y = solver.even
	}
	return
}
