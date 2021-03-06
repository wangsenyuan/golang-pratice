package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
		if !res {
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteString(strconv.Itoa(A[i]))
			buf.WriteByte(' ')
		}
		fmt.Println(buf.String())
	}
}

func solve(n int, A []int) bool {
	sort.Ints(A)
	B := make([]int, 0, n)
	var j int

	for i := 0; i < n; {
		k := i
		for k < n && A[k] == A[i] {
			k++
		}
		if k-i > 2 {
			return false
		}
		A[j] = A[i]
		j++

		if k-i == 2 {
			B = append(B, A[i])
		}
		i = k
	}
	for k := len(B) - 1; k >= 0; k-- {
		A[j] = B[k]
		if j > 0 && A[j] == A[j-1] {
			return false
		}
		j++
	}
	return true
}
