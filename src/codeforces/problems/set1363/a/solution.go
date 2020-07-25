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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, x := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(n, x, A)
		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func solve(n int, x int, A []int) bool {
	var odd int
	for i := 0; i < n; i++ {
		odd += A[i] & 1
	}

	if odd == 0 {
		return false
	}

	if x == n {
		return odd&1 == 1
	}
	// a + b = x & a is odd

	// 1, 1, 0, 0, 0, 0
	// x = 5
	if odd <= x {
		return true
	}
	// odd > x
	if x&1 == 1 {
		return true
	}
	// x is even, need at least one 0
	return odd < n
}
