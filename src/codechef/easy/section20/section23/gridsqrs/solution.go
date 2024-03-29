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
		G := make([]string, n)
		for i := 0; i < n; i++ {
			G[i], _ = reader.ReadString('\n')
		}
		res := solve1(n, G)
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
func solve(n int, G []string) int {
	U, D, L, R := make([][]int, n), make([][]int, n), make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		U[i] = make([]int, n)
		D[i] = make([]int, n)
		L[i] = make([]int, n)
		R[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if G[i][j] == '1' {
				U[i][j] = 1
				D[i][j] = 1
				L[i][j] = 1
				R[i][j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if G[i][j] == '1' {
				if i > 0 {
					U[i][j] += U[i-1][j]
				}
				if j > 0 {
					L[i][j] += L[i][j-1]
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if G[i][j] == '1' {
				if i+1 < n {
					D[i][j] += D[i+1][j]
				}
				if j+1 < n {
					R[i][j] += R[i][j+1]
				}
			}
		}
	}

	var res int

	t := NewFeeTree(n)

	dlt := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dlt[i] = make([]int, 0, 2)
	}

	for j := 0; j < n; j++ {
		t.Reset()

		r := j
		for r < n {
			for _, h := range dlt[r] {
				t.Update(h, -1)
			}
			l := r - j
			if G[l][r] == '0' {
				r++
				continue
			}
			t.Update(r, 1)
			x := r + min(R[l][r], D[l][r])
			dlt[x] = append(dlt[x], r)
			y := min(U[l][r], L[l][r])
			if r-y >= 0 {
				res += t.Get(r) - t.Get(r-y)
			} else {
				res += t.Get(r)
			}
			r++
		}
	}

	for i := 0; i < n+2; i++ {
		dlt[i] = make([]int, 0, 2)
	}

	for j := 1; j < n; j++ {
		t.Reset()
		l := j
		for l < n {
			for _, h := range dlt[l] {
				t.Update(h, -1)
			}
			r := l - j
			if G[l][r] == '0' {
				l++
				continue
			}
			t.Update(l, 1)
			x := l + min(R[l][r], D[l][r])
			dlt[x] = append(dlt[x], l)
			y := min(U[l][r], L[l][r])
			if l-y >= 0 {
				res += t.Get(l) - t.Get(l-y)
			} else {
				res += t.Get(l)
			}
			l++
		}
	}

	return res
}

type FenTree struct {
	arr []int
}

func NewFeeTree(n int) *FenTree {
	t := new(FenTree)
	t.arr = make([]int, n+1)
	return t
}

func (t *FenTree) Update(p int, v int) {
	p++
	for p < len(t.arr) {
		t.arr[p] += v
		p += p & -p
	}
}

func (t *FenTree) Get(p int) int {
	p++
	var res int
	for p > 0 {
		res += t.arr[p]
		p -= p & -p
	}
	return res
}

func (t *FenTree) GetRange(l, r int) int {
	return t.Get(r) - t.Get(l-1)
}

func (t *FenTree) Reset() {
	for i := 0; i < len(t.arr); i++ {
		t.arr[i] = 0
	}
}

type Pair struct {
	first, second int
}

func solve1(n int, G []string) int {
	sum := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		sum[i] = make([][]int, n+2)
		for j := 0; j < n+2; j++ {
			sum[i][j] = make([]int, n+2)
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if G[i-1][j-1] == '1' {
				sum[0][i][j] = 1 + sum[0][i-1][j]
				sum[1][i][j] = 1 + sum[1][i][j-1]
			}
		}
	}

	for i := n; i >= 1; i-- {
		for j := n; j >= 1; j-- {
			if G[i-1][j-1] == '1' {
				sum[2][i][j] = 1 + sum[2][i+1][j]
				sum[3][i][j] = 1 + sum[3][i][j+1]
			}
		}
	}

	f1, f2 := make([][]int, n+2), make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		f1[i] = make([]int, n+2)
		f2[i] = make([]int, n+2)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			f1[i][j] = min(sum[2][i][j], sum[3][i][j])
			f2[i][j] = min(sum[0][i][j], sum[1][i][j])
		}
	}

	var ans int

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for d := 0; f1[i][j] > d && max(i, j)+d <= n; d++ {
				if f1[i][j] > d && f2[i+d][j+d] > d {
					ans++
				}
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
