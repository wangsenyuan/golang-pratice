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
		tc--
		k := readNum(reader)
		A, B := make([]int, k), make([]int, k)
		for i := 0; i < k; i++ {
			line := readNNums(reader, 2)
			A[i] = line[0]
			B[i] = line[1]
		}
		res := solve(A, B)
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

const MOD = 1000000007
const MAX = 200010
const ROOT = 500

var phi []int

func init() {
	phi = make([]int, MAX)

	phi[1] = 1

	for i := 2; i < MAX; i++ {
		if phi[i] == 0 {
			phi[i] = i - 1
			if i < ROOT {
				for j := i * i; j < MAX; j += i {
					phi[j] = i
				}
			}
		} else {
			d := phi[i]
			j := i / d
			if j%d == 0 {
				phi[i] = d * phi[j]
			} else {
				phi[i] = (d - 1) * phi[j]
			}
		}
	}

}

func pow(a, b int) int64 {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}

func inverse(a int) int64 {
	return pow(a, MOD-2)
}

func solve(A, B []int) int {
	k := len(A)
	var P, Q int64 = 0, 1
	lo := MAX
	for i := 0; i < k; i++ {
		Q = (Q * (int64(B[i]-A[i]+1) % MOD)) % MOD
		if B[i] < lo {
			lo = B[i]
		}
	}

	for i := 1; i <= lo; i++ {
		prod := int64(phi[i])
		for j := 0; j < k; j++ {
			prod = prod * int64(B[j]/i-(A[j]-1)/i) % MOD
		}
		P += prod
		if P >= MOD {
			P -= MOD
		}
	}

	return int(P * inverse(int(MOD-Q)) % MOD)
}
