package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

// Efficient modular multiplication
func mul(a, b int) int {
	c := int64(a) * int64(b) % mod
	return int(c)
}

// Fast modular exponentiation
func pow(a, n int) int {
	ans := 1
	for n > 0 {
		if n&1 == 1 {
			ans = mul(ans, a)
		}
		a = mul(a, a)
		n >>= 1
	}
	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(a []int) int {
	root := new(Node)
	for _, num := range a {
		root.Add(num, 29)
	}
	var ans int
	var dfs func(node *Node, k int)
	dfs = func(node *Node, k int) {
		if node.children[0] != nil && node.children[1] != nil {
			i := k + 1
			ans = add(ans, mul(2*(i/2)+1, mul(node.children[0].cnt, node.children[1].cnt)))
			ans = add(ans, mul(2*((i+1)/2), mul(node.children[0].cnt, node.children[1].cnt)))
		}
		if node.children[0] == nil && node.children[1] == nil {
			i := k + 1
			ans = add(ans, mul(i, mul(node.cnt, node.cnt)))
		}
		if node.children[0] != nil {
			dfs(node.children[0], k)
		}
		if node.children[1] != nil {
			dfs(node.children[1], k+1)
		}
	}
	dfs(root, 0)

	n2 := pow(len(a), mod-2)

	ans = mul(ans, n2)
	ans = mul(ans, n2)

	return ans
}

type Node struct {
	children [2]*Node
	cnt      int
}

func (node *Node) Add(num int, pos int) {
	node.cnt++
	if pos < 0 {
		return
	}
	x := (num >> pos) & 1
	if node.children[x] == nil {
		node.children[x] = new(Node)
	}
	node.children[x].Add(num, pos-1)
}
