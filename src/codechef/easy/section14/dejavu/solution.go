package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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

func solve(A []int) int64 {
	P := make(Pairs, len(A))
	for i := 0; i < len(A); i++ {
		P[i] = Pair{A[i], i}
	}
	sort.Sort(P)
	var split int
	Q := make([]int, 0, len(P))
	var j int
	for i := 1; i <= len(P); i++ {
		if i == len(P) || P[i].second != P[i-1].second+1 {
			if i-j > split {
				split = i - j
			}
			Q = append(Q, i-j)
			j = i

		}
	}
	// apply dp on Q
	n := len(Q)
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, n)
		for ii := 0; ii < n; ii++ {
			dp[i][ii] = math.MaxInt64
		}
		// no need to merge
		dp[i][i] = 0
	}

	for j = 0; j < n; j++ {
		l := int64(Q[j])
		for i := j - 1; i >= 0; i-- {
			// dp[i][j]
			l += int64(Q[i])
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+l)
			}
		}
	}
	return dp[0][n-1] + int64(len(A)-split)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
