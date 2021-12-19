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
		res := solve(n, G)
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
