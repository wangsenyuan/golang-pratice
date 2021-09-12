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
		S, _ := reader.ReadString('\n')
		X, Y := readTwoNums(reader)
		res := solve(n, X, Y, S)
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

var D = []byte("NESW")

func solve(n int, X, Y int, S string) int64 {

	turn := func(a, b byte, Z int64, delta int) int64 {
		var i int
		for i < len(D) && D[i] != a {
			i++
		}
		var res int64
		for a != b {
			res += Z
			i = (i + delta) % 4
			a = D[i]
		}
		return res
	}

	check := func(a, b byte) int64 {
		return min(turn(a, b, int64(X), 1), turn(a, b, int64(Y), 3))
	}

	var right int64

	for i := n - 1; i >= 0; i-- {
		// try to face W
		right += check(S[i], 'W')
	}

	best := right
	var left int64

	for i := 0; i < n; i++ {
		right -= check(S[i], 'W')
		left += check(S[i], 'E')
		best = min(best, left+right)
	}

	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
