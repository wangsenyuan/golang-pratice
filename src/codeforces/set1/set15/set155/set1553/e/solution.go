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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, m)

		buf.WriteString(fmt.Sprintf("%d", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf(" %d", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(p []int, m int) []int {
	n := len(p)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		offset := i + 1 - p[i]
		if offset < 0 {
			offset += n
		}
		cnt[offset]++
	}

	var ans []int
	for i := 0; i < n; i++ {
		if cnt[i]+2*m >= n && check(p, i, m) {
			ans = append(ans, i)
		}
	}
	return ans
}

func check(p []int, k int, m int) bool {
	n := len(p)
	q := make([]int, n)
	copy(q, p[k:])
	copy(q[n-k:], p)

	return n-cycleCount(q, n) <= m
}

func cycleCount(p []int, n int) int {
	for i := 0; i < len(p); i++ {
		p[i]--
	}
	vis := make([]bool, n)
	var ans int

	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		j := i
		for !vis[j] {
			vis[j] = true
			j = p[j]
		}
		ans++
	}
	return ans
}
