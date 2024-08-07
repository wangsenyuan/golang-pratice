package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		a, b, c := readThreeNums(reader)
		res := solve(a, b, c)
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

func solve(a, b, c int) int {
	if a+1 != c {
		return -1
	}
	h := bits.Len(uint(a))

	b -= (1 << h) - (a + 1)
	if b > 0 {
		h += (b + a) / (a + 1)
	}
	return h
}
func solve2(a, b, c int) int {
	if a+1 != c {
		return -1
	}
	if a == 0 {
		return b
	}
	h := 31 - bits.LeadingZeros32(uint32(a))
	left := (1 << (h + 1)) - (a + 1)
	if left < b {
		b -= left
		base := a - ((1 << h) - 1)
		base = base*2 + left
		h += (b + base - 1) / base
	}
	h++
	return h
}

func solve1(a, b, c int) int {
	if a+1 != c {
		return -1
	}
	if a+b+c == 1 {
		return 0
	}
	// 分层
	cur := 1
	var next int
	res := 1
	for i := 0; i < a+b; i++ {
		if cur == 0 {
			cur, next = next, cur
			res++
		}
		cur--
		next++
		if i < a {
			next++
		}
	}
	return res
}
