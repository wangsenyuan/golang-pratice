package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	A := readString(reader)
	B := readString(reader)
	res := solve(n, A, B)
	fmt.Printf("%d %d\n", res[0], res[1])
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(n int, A string, B string) []int {
	// loop until (i,j) == 0, 0, and record its position, p
	// then n / p * x (x is the in-eq in prefix p)
	// and loop the remaing part
	i, j := 0, 0
	var p int
	la := len(A)
	lb := len(B)
	res := []int{0, 0}
	for {
		if A[i] != B[j] {
			if beats(A[i], B[j]) {
				res[1]++
			} else {
				res[0]++
			}
		}
		p++
		i = (i + 1) % la
		j = (j + 1) % lb
		if i == 0 && j == 0 {
			break
		}
	}

	res[0] *= n / p
	res[1] *= n / p
	r := n % p
	for r > 0 {
		if A[i] != B[j] {
			if beats(A[i], B[j]) {
				res[1]++
			} else {
				res[0]++
			}
		}
		r--
		i = (i + 1) % la
		j = (j + 1) % lb
	}
	return res
}

func beats(x, y byte) bool {
	if x == 'R' {
		return y == 'S'
	}
	if x == 'S' {
		return y == 'P'
	}
	if x == 'P' {
		return y == 'R'
	}
	return false
}