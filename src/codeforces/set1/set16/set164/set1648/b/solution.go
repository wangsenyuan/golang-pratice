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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, _ := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if res {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A []int) bool {
	sort.Ints(A)

	var n int
	for i := 1; i <= len(A); i++ {
		if i == len(A) || A[i] > A[i-1] {
			A[n] = A[i-1]
			n++
		}
	}
	if A[0] != 1 {
		return false
	}
	var R []int
	for i, j := 1, 0; i < A[n-1]; i++ {
		for j < n && A[j] < i {
			j++
		}
		if j < n && A[j] == i {
			j++
		} else {
			R = append(R, i)
		}
	}

	c := A[n-1]

	cnt := make([]int, c+1)
	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}

	for i := 2; i <= c; i++ {
		cnt[i] += cnt[i-1]
	}

	for _, r := range R {
		for i := 0; i < n; i++ {
			y := A[i]
			if int64(r)*int64(y) > int64(c) {
				break
			}
			a := r * y
			b := c + 1
			if int64(r+1)*int64(y) <= int64(c) {
				b = (r + 1) * y
			}
			if cnt[b-1] > cnt[a-1] {
				return false
			}
		}
	}
	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
