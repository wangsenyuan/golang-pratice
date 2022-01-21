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
		x, s := readTwoNums(reader)
		res := solve(x, s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const INF = 1000000001

func solve(x int, s int) int {

	s -= x

	cnt := make([]int, 30)

	check := func(m int) bool {
		for i := 0; i < 30; i++ {
			cnt[i] = (s >> uint(i)) & 1
		}

		for i := 29; i > 0; i-- {
			if (x>>uint(i))&1 == 1 {
				if cnt[i] > m {
					cnt[i-1] += 2 * (cnt[i] - m)
				}
			} else {
				cnt[i-1] += 2 * cnt[i]
			}
		}

		if x&1 == 1 {
			return cnt[0] <= m
		}
		return cnt[0] == 0
	}

	l, r := 0, INF

	for l < r {
		mid := (l + r) / 2
		if check(mid - 1) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if r >= INF {
		return -1
	}

	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
