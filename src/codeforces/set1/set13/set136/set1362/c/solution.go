package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		var num uint64
		s, _ := reader.ReadBytes('\n')
		readUint64(s, 0, &num)

		fmt.Println(solve(num))
	}
}

const N = 64

var dp []uint64

func init() {
	dp = make([]uint64, N)
	dp[0] = 1

	for i := 1; i < N; i++ {
		dp[i] = (dp[i-1] << 1) + uint64(i+1)
	}
}

func solve1(num uint64) uint64 {
	var res uint64

	for i := N - 1; i > 0; i-- {
		if (num>>i)&1 == 1 {
			res += dp[i-1] + uint64(i+1)
		}
	}

	if num&1 == 1 {
		res++
	}

	return res
}

func solve(num uint64) uint64 {
	res := num

	for num > 0 {
		res += num / 2
		num >>= 1
	}

	return res
}
