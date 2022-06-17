package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	menus := make([]string, n)
	for i := 0; i < n; i++ {
		menus[i] = readString(reader)
	}
	res := solve(menus)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(menus []string) int {
	var res int

	for _, cur := range menus {
		res += check(cur)
	}

	return res
}

func check(menu string) int {
	cnt := make([]int, 3)
	i := len(menu) - 1
	if menu[i] == ' ' {
		i--
	}
	for i > 0 && isDigit(menu[i]) {
		if menu[i] == '8' {
			cnt[0]++
		} else if menu[i] == '5' {
			cnt[1]++
		} else if menu[i] == '3' {
			cnt[2]++
		} else {
			return 0
		}
		i--
	}
	if cnt[0] >= cnt[1] && cnt[1] >= cnt[2] {
		return 1
	}
	return 0
}

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}
