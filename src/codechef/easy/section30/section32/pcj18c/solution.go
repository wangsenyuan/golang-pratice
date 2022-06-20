package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, x, k := readThreeNums(reader)
		a, b := solve(n, x, k)
		buf.WriteString(fmt.Sprintf("%d %d\n", a, b))
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
		if s[i] == '\n' {
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

func solve(n int, a int, k int) (int64, int64) {
	if k == 1 {
		return int64(a), 1
	}
	//
	N := int64(n)
	sum := (N - 2) * 180
	A := int64(a)
	K := int64(k)
	// sum = (A1 + An) * n / 2
	// An = 2 * sum / n - A1
	// d := (An - A1) / (n - 1)
	// A[k] = A[1] + (k - 1) * d
	// A[k] = A + (k - 1) * (2 * sum / n - A - A) / (n - 1)
	// A + (k - 1) * 2 * (sum / n - A)/ (n - 1)
	//( (n - 1) * A + (k - 1) * 2 * (sum / n  - A)) / (n - 1)
	// (n - 1) * A * n + (k - 1) * 2 * (sum - n * A)
	P := (N-1)*N*A + (K-1)*2*(sum-N*A)
	Q := N * (N - 1)
	g := gcd(P, Q)

	P /= g
	Q /= g
	return P, Q
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
