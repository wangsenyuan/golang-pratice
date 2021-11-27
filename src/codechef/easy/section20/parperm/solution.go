package main

import (
	"bufio"
	"bytes"
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
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		res := solve(n, k)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < k; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
}

const MAX_N = 100001

var lpf []int

func init() {
	set := make([]bool, MAX_N)
	lpf = make([]int, MAX_N)
	for i := 2; i < MAX_N; i++ {
		if !set[i] {
			for j := i; j < MAX_N; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
				set[j] = true
			}
		}
	}
}

func solve(n int, K int) []int {
	A := make([]int, 0, n)
	B := make([]int, 0, n)

	A = append(A, 1)

	for i := 2; i <= n/2; i++ {
		B = append(B, i)
	}
	for i := n/2 + 1; i <= n; i++ {
		if lpf[i] == i {
			A = append(A, i)
		} else {
			B = append(B, i)
		}
	}

	if K < n-len(A) && n-K < n-len(A) {
		return nil
	}
	p := len(A)
	for len(B) != K && K != p {
		B = append(B, A[p-1])
		p--
	}
	if len(B) == K {
		return B
	}
	return A[:p]
}