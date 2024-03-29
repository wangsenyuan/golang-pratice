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
		n := readNum(reader)
		A := readNNums(reader, n)
		B := solve(A)

		if len(B) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < len(B); i++ {
				buf.WriteString(fmt.Sprintf("%d ", B[i]))
			}
			buf.WriteByte('\n')
		}
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

func solve(A []int) []int {
	n := len(A)

	if n == 1 {
		return A
	}
	sort.Ints(A)
	// 如果存在一个数的个数多于2个，则没有答案
	for i := 0; i < n; i++ {
		cnt := 1
		if i > 0 && A[i] == A[i-1] {
			cnt++
		}
		if i+1 < n && A[i] == A[i+1] {
			cnt++
		}
		if cnt == 3 {
			return nil
		}
	}
	if A[n-1] == A[n-2] {
		return nil
	}
	B := make([]int, n)
	front, end := 0, n-1

	for i := 0; i < n; {
		B[end] = A[i]
		end--
		if i+1 < n && A[i] == A[i+1] {
			B[front] = A[i+1]
			front++
			i += 2
		} else {
			i++
		}
	}
	return B
}
