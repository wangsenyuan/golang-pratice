package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	for tc > 0 {
		tc--

		s, _ := reader.ReadString('\n')

		fmt.Println(solve(s[:len(s)-1]))
	}
}

func solve(s string) int64 {
	n := len(s)
	var sum int64

	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			sum++
		}
	}
	N := int64(n)

	res := sum * N * (N + 1) / 2

	// flip the whole string
	// res += sum
	// fix L at 0

	for i := 1; i < n; i++ {
		// if we fix i as L
		if s[i] == s[i-1] {
			res -= N - int64(i)
		} else {
			res += N - int64(i)
		}
	}
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			res -= int64(i) + 1
		} else {
			res += int64(i) + 1
		}
	}

	return res
}