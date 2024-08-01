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
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		c := readNNums(reader, n)
		res := solve(a, b, c)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
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

func solve(a, b, c []int) []int {
	tot := sum(a)
	exp := (tot + 2) / 3
	n := len(a)
	check := func(x, y, z []int) (bool, []int) {
		// x for first, z for last, y for mid
		var first, mid, last int
		var l int
		for l < n && first < exp {
			first += x[l]
			l++
		}
		r := n - 1
		for r > 0 && last < exp {
			last += z[r]
			r--
		}
		for i := l; i <= r; i++ {
			mid += y[i]
		}
		if mid >= exp {
			return true, []int{l, r + 1}
		}
		return false, nil
	}

	arr := [][]int{a, b, c}
	ans := make([][]int, 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			k := 3 - i - j
			ok, res := check(arr[i], arr[j], arr[k])
			if !ok {
				continue
			}
			l, r := res[0], res[1]
			ans[i] = []int{1, l}
			ans[j] = []int{l + 1, r}
			ans[k] = []int{r + 1, n}
			return concate(ans[0], ans[1], ans[2])
		}
	}

	return nil
}

func concate(a, b, c []int) []int {
	return []int{a[0], a[1], b[0], b[1], c[0], c[1]}
}

func sum(arr []int) int {
	var res int
	for _, num := range arr {
		res += num
	}
	return res
}
