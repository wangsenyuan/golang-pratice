package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		scanner.Scan()
		A := scanner.Text()
		fmt.Println(solve(n, k, A))
	}
}

func solve(n, k int, A string) int {
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if A[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = 1
			if i > 0 {
				dp[i] += dp[i-1]
			}
		}
	}

	fp := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		if A[i] == '0' {
			fp[i] = 0
		} else {
			fp[i] = 1
			if i < n-1 {
				fp[i] += fp[i+1]
			}
		}
	}

	res := k + fp[k]

	for i := k; i < n-1; i++ {
		a := dp[i-k] + k + fp[i+1]
		res = max(res, a)
	}

	res = max(res, k+dp[n-k])

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
