package main

import (
	"bufio"
	"bytes"
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

	k := readNum(reader)
	arrs := make([][]int, k)
	for i := 0; i < k; i++ {
		n := readNum(reader)
		arrs[i] = readNNums(reader, n)
	}

	res := solve(arrs)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, row := range res {
		for i := 0; i < 2; i++ {
			buf.WriteString(fmt.Sprintf("%d ", row[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func solve(arrs [][]int) [][]int {
	// sum - i
	records := make(map[int]Pair)
	for i, arr := range arrs {
		var sum int
		for _, num := range arr {
			sum += num
		}

		for j := 0; j < len(arr); j++ {
			tmp := sum - arr[j]
			if v, ok := records[tmp]; ok {
				return [][]int{
					{v.first + 1, v.second + 1},
					{i + 1, j + 1},
				}
			}
		}

		for j := 0; j < len(arr); j++ {
			tmp := sum - arr[j]
			records[tmp] = Pair{i, j}
		}
	}

	return nil
}

type Pair struct {
	first  int
	second int
}
