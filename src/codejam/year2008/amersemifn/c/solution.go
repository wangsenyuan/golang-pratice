package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func readNFloat64s(scanner *bufio.Scanner, n int) []float64 {
	res := make([]float64, n)

	pos := 0
	scanner.Scan()
	bs := scanner.Bytes()
	//fmt.Printf("[debug] %s\n", string(bs))
	for i := 0; i < n; i++ {
		for pos < len(bs) && bs[pos] == ' ' {
			pos++
		}

		pos = readFloat64(bs, pos, &res[i])
	}
	return res
}

func main() {
	f, err := os.Open("./C-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		M, Q := readTwoNums(scanner)
		qs := make([][]float64, Q)
		for j := 0; j < Q; j++ {
			qs[j] = readNFloat64s(scanner, 4)
		}
		res := solve(M, Q, qs)
		fmt.Printf("Case #%d: %.7f\n", i, res)
	}
}

func solve(M int, Q int, questions [][]float64) float64 {
	cur := make([]float64, 0, 4)
	for i := 0; i < 4; i++ {
		if questions[0][i] > 0 {
			cur = append(cur, questions[0][i])
		}
		reverseSort(cur)
	}

	for i := 1; i < Q; i++ {
		q := questions[i]
		vs := make([]float64, 0, len(cur))

		for _, u := range cur {
			for j := 0; j < 4; j++ {
				if q[j] > 0 {
					vs = append(vs, u*q[j])
				}
			}
		}
		reverseSort(vs)
		if M < len(vs) {
			cur = vs[0:M]
		} else {
			cur = vs
		}
	}
	var ans float64

	for i := 0; i < M && i < len(cur); i++ {
		ans += cur[i]
	}

	return ans
}

func reverseSort(arr []float64) {
	sort.Float64s(arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
