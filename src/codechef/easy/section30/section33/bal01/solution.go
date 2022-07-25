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
		s := readString(reader)
		res := solve(s[:n])
		buf.WriteString(fmt.Sprintf("%s\n", res))
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(s string) string {
	cnt := make([]int, 2)
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			cnt[0]++
		} else if s[i] == '1' {
			cnt[1]++
		}
	}
	que := n - cnt[0] - cnt[1]
	best := n
	// (cnt[0] + a) - (cnt[1] + que - a) 最小
	for a := 0; a <= que; a++ {
		x := cnt[0] + a
		y := cnt[1] + que - a
		best = min(best, abs(x-y))
	}

	for a := 0; a <= que; a++ {
		x := cnt[0] + a
		y := cnt[1] + que - a
		tmp := abs(x - y)
		if tmp == best {
			return replace(s, a)
		}
	}

	return ""
}

func replace(s string, a int) string {
	buf := []byte(s)

	var i int
	for i < len(buf) && a > 0 {
		if buf[i] == '?' {
			buf[i] = '0'
			a--
		}
		i++
	}

	for i < len(buf) {
		if buf[i] == '?' {
			buf[i] = '1'
		}
		i++
	}

	return string(buf)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
