package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n, m := readTwoNums(scanner)

	A := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		A[i] = scanner.Text()
	}
	Q := readNum(scanner)
	quries := readNNums(scanner, Q)
	res := solve(n, m, A, Q, quries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(n, m int, A []string, Q int, quries []int) []int {
	// dp[i][j][0][L] means flips to have width-L, starting at top-left (i, j), color white
	L := min(n, m)

	xp := make([][][][]int, n)
	yp := make([][][][]int, n)
	dp := make([][][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][][]int, m)
		xp[i] = make([][][]int, m)
		yp[i] = make([][][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([][]int, 2)
			xp[i][j] = make([][]int, 2)
			yp[i][j] = make([][]int, 2)
			for k := 0; k < 2; k++ {
				dp[i][j][k] = make([]int, L+1)
				xp[i][j][k] = make([]int, L+1)
				yp[i][j][k] = make([]int, L+1)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == '1' {
				dp[i][j][0][1] = 1
				dp[i][j][1][1] = 0
				xp[i][j][0][1] = 1
				xp[i][j][1][1] = 0
				yp[i][j][0][1] = 1
				yp[i][j][1][1] = 0
			} else {
				dp[i][j][1][1] = 1
				dp[i][j][0][1] = 0
				xp[i][j][1][1] = 1
				xp[i][j][0][1] = 0
				yp[i][j][1][1] = 1
				yp[i][j][0][1] = 0
			}
		}
	}

	for l := 2; l <= L; l++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				for color := 0; color < 2; color++ {
					xp[i][j][color][l] = xp[i][j][color][l-1]
					yp[i][j][color][l] = yp[i][j][color][l-1]

					if l%2 == 0 {
						// different color
						if j+l <= m && int(A[i][j+l-1]-'0') != 1-color {
							// flip
							xp[i][j][color][l]++
						}
						if i+l <= n && int(A[i+l-1][j]-'0') != 1-color {
							yp[i][j][color][l]++
						}
					} else {
						// same
						if j+l <= m && int(A[i][j+l-1]-'0') != color {
							// flip
							xp[i][j][color][l]++
						}
						if i+l <= n && int(A[i+l-1][j]-'0') != color {
							yp[i][j][color][l]++
						}
					}
				}
			}
		}
	}

	fp := make([]int, L+1)
	for i := 0; i <= L; i++ {
		fp[i] = math.MaxInt32
	}
	fp[0] = 0
	fp[1] = 0

	for l := 2; l <= L; l++ {
		for i := 0; i+l <= n; i++ {
			for j := 0; j+l <= m; j++ {
				for color := 0; color < 2; color++ {
					dp[i][j][color][l] = dp[i][j][color][l-1]
					k := color
					if l%2 == 0 {
						k = 1 - color
					}
					dp[i][j][color][l] += xp[i+l-1][j][k][l-1]
					dp[i][j][color][l] += yp[i][j+l-1][k][l-1]
					if int(A[i+l-1][j+l-1]-'0') != color {
						dp[i][j][color][l]++
					}
					if fp[l] > dp[i][j][color][l] {
						fp[l] = dp[i][j][color][l]
					}
				}
			}
		}
	}

	ans := make([]int, Q)

	find := func(flips int) int {
		i := sort.Search(L+1, func(i int) bool {
			return fp[i] > flips
		})
		return i - 1
	}

	for i := 0; i < Q; i++ {
		ans[i] = find(quries[i])
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
