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
		S := readString(reader)
		T := readString(reader)
		res := solve(n, S, T)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, S string, T string) int {
	cnt := make([]int, 4)
	for i := 0; i < n; i++ {
		var x int
		if S[i] == '1' {
			x += 2
		}
		if T[i] == '1' {
			x += 1
		}
		cnt[x]++
	}

	var res int
	x := min(cnt[1], cnt[2])
	res += x

	cnt[1] -= x
	cnt[2] -= x

	if cnt[1] > 0 {
		res += min(cnt[1], cnt[3])
		cnt[3] -= min(cnt[1], cnt[3])
	}
	if cnt[2] > 0 {
		res += min(cnt[2], cnt[3])
		cnt[3] -= min(cnt[2], cnt[3])
	}

	if cnt[0] > 0 {
		res += min(cnt[0], cnt[3])
		cnt[3] -= min(cnt[0], cnt[3])
	}

	res += cnt[3] / 2

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
