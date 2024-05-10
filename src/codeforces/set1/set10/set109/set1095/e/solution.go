package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)

	s := readString(reader)

	res := solve(s)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(s string) int {
	n := len(s)
	if n%2 == 1 {
		return 0
	}

	fs := flip(reverse(s))
	pref := make([]int, n+1)
	suf := make([]int, n+1)
	lok := make([]bool, n+1)
	rok := make([]bool, n+1)
	lok[0] = true
	rok[n] = true

	for i := 0; i < n; i++ {
		pref[i+1] = pref[i]
		if s[i] == '(' {
			pref[i+1]++
		} else {
			pref[i+1]--
		}
		lok[i+1] = lok[i] && pref[i+1] >= 0

		suf[n-i-1] = suf[n-i]
		if fs[i] == '(' {
			suf[n-i-1]++
		} else {
			suf[n-i-1]--
		}

		rok[n-i-1] = rok[n-i] && suf[n-i-1] >= 0
	}

	var ans int

	for i := 0; i < n; i++ {
		if !lok[i] || !rok[i+1] {
			continue
		}

		if s[i] == '(' {
			if pref[i] > 0 && pref[i]-1-suf[i+1] == 0 {
				ans++
			}
		} else {
			if pref[i]+1-suf[i+1] == 0 {
				ans++
			}
		}
	}

	return ans
}

func flip(s string) string {
	buf := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			buf[i] = ')'
		} else {
			buf[i] = '('
		}
	}
	return string(buf)
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
