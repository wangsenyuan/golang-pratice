package main

import (
	"bufio"
	"bytes"
	"fmt"
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

func solve(A []int) int {
	var sum int
	same := true
	for _, num := range A {
		sum += num
		if num != A[0] {
			same = false
		}
	}
	if same {
		return 0
	}
	// sum <= 1e8
	n := len(A)
	// 如果最后的长度是m, 则每个数字是 sum / m

	check := func(x int) bool {
		for i := 0; i < n; {
			var cur int
			for i < n && cur+A[i] <= x {
				cur += A[i]
				i++
			}
			if cur < x {
				return false
			}
		}
		return true
	}

	var factors []int

	for i := 2; i <= sum/i; i++ {
		if sum%i == 0 {
			if i < n {
				factors = append(factors, i)
			}
			j := sum / i
			if j < n && i != j {
				factors = append(factors, j)
			}
		}
	}

	sort.Ints(factors)

	for i := len(factors) - 1; i >= 0; i-- {
		m := factors[i]
		if check(sum / m) {
			return n - m
		}
	}

	return n - 1
}
