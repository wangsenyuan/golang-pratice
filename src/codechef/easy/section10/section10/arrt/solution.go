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
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		C := solve(n, A, B)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", C[i]))
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

func solve(n int, A []int, B []int) []int {
	min := 2 * n

	for i := 0; i < n; i++ {
		tmp := (A[0] + B[i]) % n
		if tmp < min {
			min = tmp
		}
	}

	C := make([]int, n)
	D := make([]int, n)

	var cnt int
	for i := 0; i < n; i++ {
		tmp := (A[0] + B[i]) % n
		if tmp == min {
			cnt++
			for j := 0; j < n; j++ {
				cur := (A[j] + B[(i+j)%n]) % n
				if cnt == 1 {
					C[j] = cur
				} else {
					D[j] = cur
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if C[i] < D[i] {
			return C
		}
		if C[i] > D[i] {
			return D
		}
	}

	return C
}
