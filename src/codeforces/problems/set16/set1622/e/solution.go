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
		n, m := readTwoNums(reader)
		X := readNNums(reader, n)
		A := make([]string, n)
		for i := 0; i < n; i++ {
			A[i], _ = reader.ReadString('\n')
		}
		res := solve(n, m, X, A)
		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func solve(n int, m int, X []int, A []string) []int {

	var ans int = -1
	best := make([]int, m)

	for mask := 0; mask < 1<<n; mask++ {

		val := make([]int, m)

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if A[i][j] == '1' {
					if (mask>>i)&1 == 1 {
						val[j]++
					} else {
						val[j]--
					}
				}
			}
		}

		p := make([]int, m)
		for i := 0; i < m; i++ {
			p[i] = i
		}
		sort.Slice(p, func(i, j int) bool {
			return val[p[i]] < val[p[j]]
		})

		var tmp int
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				tmp -= X[i]
			} else {
				tmp += X[i]
			}
		}

		for i := 0; i < m; i++ {
			tmp += val[p[i]] * (i + 1)
		}

		if tmp > ans {
			ans = tmp
			copy(best, p)
		}
	}
	res := make([]int, m)
	for i := 0; i < m; i++ {
		res[best[i]] = i + 1
	}

	return res
}
