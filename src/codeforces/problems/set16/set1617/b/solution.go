package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a, b, c := solve(n)
		buf.WriteString(fmt.Sprintf("%d %d %d\n", a, b, c))
	}
	fmt.Print(buf.String())
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int) (a int, b int, c int) {
	// a = x * c
	// b = y * c
	// gcd(x, y) = 1
	// (x + y + 1) * c = n
	// a != b != c
	// if c == 1
	// then x + y + 1 = n
	// x + y = n - 1
	c = 1
	if n%2 == 0 {
		// n is even, and  n - 1 is odd
		a = (n - 1) / 2
		b = n - 1 - a
		return
	}
	// n % 2 == 1, if c == 1
	// then a + b + 1 = 2 * x + 1
	// a + b = 2 * x
	// a is odd, and b is odd, find gcd(a, b) = 1
	// if x is even, then a = x - 1, b = x + 1
	// else a = x - 2, b = x + 2
	x := n / 2
	if x%2 == 0 {
		a = x - 1
		b = x + 1
	} else {
		a = x - 2
		b = x + 2
	}
	return
}
