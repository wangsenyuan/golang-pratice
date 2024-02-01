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
		res := solve(a)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

const D = 31

func solve(a []int) bool {
	var sum int
	for _, num := range a {
		sum += num
	}

	n := len(a)

	if sum%n != 0 {
		return false
	}
	avg := sum / n

	var free int
	cnt := make([]int, D)
	for i := 0; i < n; i++ {
		diff := avg - a[i]
		if diff == 0 {
			free++
			continue
		}
		receive, give := calc(abs(diff))
		if receive < 0 {
			return false
		}
		if diff < 0 {
			receive, give = give, receive
		}
		cnt[receive]++
		cnt[give]--
	}
	if free == n {
		return true
	}

	for i := 0; i < D; i++ {
		if cnt[i] != 0 {
			return false
		}
	}

	return true
}

func solve1(a []int) bool {
	var sum int
	for _, num := range a {
		sum += num
	}

	n := len(a)

	if sum%n != 0 {
		return false
	}
	avg := sum / n

	var free int

	nodes := make([][]Node, D)

	for i := 0; i < n; i++ {
		diff := avg - a[i]
		if diff == 0 {
			free++
			continue
		}
		receive, give := calc(abs(diff))
		if receive < 0 {
			return false
		}
		if diff < 0 {
			receive, give = give, receive
		}
		if nodes[receive] == nil {
			nodes[receive] = make([]Node, 0, 1)
		}
		nodes[receive] = append(nodes[receive], Node{i, receive, give})
	}
	if free == n {
		return true
	}

	popBack := func(x int) Node {
		m := len(nodes[x])
		res := nodes[x][m-1]
		nodes[x] = nodes[x][:m-1]
		return res
	}

	for d := D - 1; d >= 0; d-- {
		if len(nodes[d]) == 0 {
			continue
		}

		for _, node := range nodes[d] {
			expect := node.receive
			x := node.give
			for x != expect {
				if len(nodes[x]) == 0 {
					return false
				}
				last := popBack(x)
				x = last.give
			}
		}
	}

	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func calc(num int) (receive int, give int) {
	var lo int
	for (num>>lo)&1 == 0 {
		lo++
	}
	// num >> lo & 1 == 1
	hi := lo
	for (num>>hi)&1 == 1 {
		num ^= 1 << hi
		hi++
	}
	// num >> hi & 1 == 0
	if num != 0 {
		return -1, -1
	}
	return hi, lo
}

type Node struct {
	id      int
	receive int
	give    int
}
