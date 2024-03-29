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
		n := readNum(reader)
		nums := readNNums(reader, n)
		res := solve(nums, tc)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
		tc--
	}
	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const X = 100_010

var primes []int
var lpf []int
var ind [X]int

func init() {
	lpf = make([]int, X)
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if primes[j]*i >= X {
				break
			}
			lpf[primes[j]*i] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}
}

func solve(a []int, tc int) bool {
	set := make(map[int]bool)
	check := func(x int) bool {
		if x == 1 {
			return false
		}

		if x < X && lpf[x] == x {
			// a prime number
			if ind[x] == tc {
				return true
			}
			ind[x] = tc
			return false
		}
		for i := 0; i < len(primes) && primes[i]*primes[i] <= x; i++ {
			if x%primes[i] == 0 {
				if ind[primes[i]] == tc {
					return true
				}
				ind[primes[i]] = tc
				for x%primes[i] == 0 {
					x /= primes[i]
				}
			}
		}
		if x == 1 {
			return false
		}
		if x < X {
			if ind[x] == tc {
				return true
			}
			ind[x] = tc
		} else {
			// x >= X
			if set[x] {
				return true
			}
			set[x] = true
		}
		return false
	}

	for _, x := range a {
		if check(x) {
			return true
		}
	}

	return false
}
