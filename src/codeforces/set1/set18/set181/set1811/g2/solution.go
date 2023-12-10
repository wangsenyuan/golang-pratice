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
		n, k := readTwoNums(reader)
		p := readNNums(reader, n)
		res := solve(p, k)
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

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func nCr(n int, r int) int {
	if r < 0 || n < r {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

const N = 5010

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}

	I[N-1] = pow(F[N-1], mod-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func solve(p []int, k int) int {
	// 首先要找到最长的m m % k = 0
	// 这个可以用贪心吗？
	// dp[i] 表示到i为止（且都是连续的block的数量）
	// dp[i] = dp[i-1] or 选出k个和 p[i]相同的tile后的最大值 + 1
	// 然后再计数
	// fp[i][j] 表示到i为止有j个block的数量 fp[i][j] = fp[i][j-1] + 使用i个连续block的数量

	if k == 1 {
		return 1
	}

	n := len(p)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = []int{0, 0}
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		cnt := 1
		for j := i - 1; j > 0; j-- {
			if p[j-1] == p[i-1] {
				cnt++

				if cnt == k {
					dp[i][1] = dp[j-1][1] + 1
				}
				if cnt >= k {
					if dp[j-1][1]+1 < dp[i][1] {
						break
					}
					dp[i][0] = add(dp[i][0], mul(dp[j-1][0], nCr(cnt-2, k-2)))
				}
			}
		}
		if dp[i][1] < dp[i-1][1] {
			dp[i][0] = 0
			dp[i][1] = dp[i-1][1]
		}
		if dp[i][1] == dp[i-1][1] {
			dp[i][0] = add(dp[i][0], dp[i-1][0])
		}
	}

	return dp[n][0]
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
