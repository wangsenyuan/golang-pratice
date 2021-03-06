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

	n := readNum(reader)

	A := readNNums(reader, n)

	cnt := solve(n, A)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", cnt))

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", A[i]))
	}
	buf.WriteByte('\n')
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

func solve(n int, A []int) int {
	sort.Ints(A)

	B := make([]int, n)

	var j int
	for i := 1; i < n; i += 2 {
		B[i] = A[j]
		j++
	}

	k := n - 1
	if k&1 == 1 {
		// B[k] is already filled
		k--
	}

	for i := 0; i <= k; {
		cur := A[j]
		j++
		var ok = true
		if i > 0 && cur == B[i-1] {
			ok = false
		}
		if i+1 < n && cur == B[i+1] {
			ok = false
		}
		if ok {
			B[i] = cur
			i += 2
			continue
		}
		// i not filled
		B[k] = cur
		k -= 2
	}
	var cnt int

	for i := 1; i < n; i += 2 {
		if B[i] < B[i-1] && i+1 < n && B[i] < B[i+1] {
			cnt++
		}
	}
	copy(A, B)

	return cnt
}
