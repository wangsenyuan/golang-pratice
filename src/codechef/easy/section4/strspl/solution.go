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
		s, _ := reader.ReadString('\n')
		s = s[:len(s)-1]
		k, m, res := solve(s)
		buf.WriteString(fmt.Sprintf("%d %d\n", k, m))
		for i := 0; i < len(s); i++ {
			if res[i] == 0 {
				buf.WriteString(fmt.Sprintf("%d ", i+1))
			}
		}
		buf.WriteByte('\n')
		for i := 0; i < len(s); i++ {
			if res[i] == 1 {
				buf.WriteString(fmt.Sprintf("%d ", i+1))
			}
		}
		buf.WriteByte('\n')
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

func solve(s string) (int, int, []int) {
	n := len(s)
	res := make([]int, n)
	var m int
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			res[i] = 0
			res[j] = 0
		} else {
			res[i] = 1
			res[j] = 1
			m += 2
		}
	}
	return n - m, m, res
}
