package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := readString(reader)
	b := readString(reader)
	res := solve(a, b)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func solve(s string, t string) [][]int {
	n := len(s)
	m := len(t)

	s = "$" + s + "$"

	pos1 := make([]bool, len(s))
	pos2 := make([]bool, len(s))

	t += "#"

	var ans [][]int
	for i := 0; i < m; {
		var found bool
		cur := Pair{-1, -1}
		for j := 1; j <= n; j++ {
			pos1[j] = s[j] == t[i]
			pos2[j] = s[j] == t[i]
			found = found || pos1[j]
		}
		for ln := 1; found; i, ln = i+1, ln+1 {
			found = false
			for j := 1; j <= n; j++ {
				if pos2[j] {
					cur = Pair{j + ln - 1, j}
				}
				pos2[j] = pos2[j+1] && s[j] == t[i+1]
				found = found || pos2[j]
			}
			for j := n; j >= 1; j-- {
				if pos1[j] {
					cur = Pair{j - ln + 1, j}
				}
				pos1[j] = pos1[j-1] && s[j] == t[i+1]
				found = found || pos1[j]
			}
		}
		if cur.first < 0 {
			return nil
		}
		ans = append(ans, []int{cur.first, cur.second})
	}

	return ans
}
func solve1(s string, t string) [][]int {
	// dp[i] 表示到t[i]为止能够获得的最少的操作数
	// dp[i] = dp[j] + 1 if t[j+1...i]是s的一个子串（反字串）
	n := len(s)
	fr := NewTrie()
	bk := NewTrie()

	for i := 0; i < n; i++ {
		var cur int
		for j := i; j < n; j++ {
			x := int(s[j] - 'a')
			cur = fr.Add(cur, x, Pair{i, j})
		}
		cur = 0
		for j := i; j >= 0; j-- {
			x := int(s[j] - 'a')
			cur = bk.Add(cur, x, Pair{i, j})
		}
	}

	m := len(t)
	dp := make([][]int, m)

	update := func(i int, j int, l, r int) {
		if j < 0 {
			dp[i] = []int{1, -1, l, r}
			return
		}
		v := dp[j]
		if v[0] < 0 {
			return
		}
		if dp[i][0] < 0 || v[0]+1 < dp[i][0] {
			dp[i] = []int{v[0] + 1, j, l, r}
		}
	}

	for i := 0; i < m; i++ {
		dp[i] = []int{-1, -1, -1, -1}
		var cur int
		for j := i; j >= 0; j-- {
			if i-j+1 > n {
				// too much
				break
			}
			x := int(t[j] - 'a')
			cur = fr.next[cur][x]
			if cur == 0 {
				break
			}
			v := fr.val[cur]
			update(i, i-(v.second-v.first+1), v.second+1, v.first+1)
		}
		cur = 0
		for j := i; j >= 0; j-- {
			if i-j+1 > n {
				// too much
				break
			}
			x := int(t[j] - 'a')
			cur = bk.next[cur][x]
			if cur == 0 {
				break
			}
			v := bk.val[cur]
			update(i, i-(v.first-v.second+1), v.second+1, v.first+1)
		}
	}

	if dp[m-1][0] < 0 {
		// no answer
		return nil
	}

	var res [][]int

	for i := m - 1; i >= 0; {
		v := dp[i]
		res = append(res, []int{v[2], v[3]})
		i = v[1]
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type Trie struct {
	next [][]int
	val  []Pair
}

func NewTrie() *Trie {
	next := make([][]int, 1)
	next[0] = make([]int, 26)
	val := make([]Pair, 1)
	return &Trie{next, val}
}

func (tr *Trie) Add(node int, x int, v Pair) int {
	if tr.next[node][x] == 0 {
		tr.next = append(tr.next, make([]int, 26))
		tr.val = append(tr.val, v)
		tr.next[node][x] = len(tr.next) - 1
	}
	return tr.next[node][x]
}
