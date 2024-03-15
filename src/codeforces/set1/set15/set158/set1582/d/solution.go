package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(fmt.Sprintf("%s\n", s[1:len(s)-1]))
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

func solve(a []int) []int {
	n := len(a)
	b := make([]int, n)

	if n%2 == 0 {
		for i := 0; i < n; i += 2 {
			b[i] = a[i+1]
			b[i+1] = -a[i]
		}
	} else {
		for i := 0; i < n-3; i += 2 {
			b[i] = a[i+1]
			b[i+1] = -a[i]
		}
		i, j, k := n-3, n-2, n-1
		if a[i]+a[j] == 0 {
			j, k = k, j
		}
		if a[i]+a[j] == 0 {
			i, k = k, i
		}
		b[i] = -a[k]
		b[j] = -a[k]
		b[k] = a[i] + a[j]
	}
	return b
}
