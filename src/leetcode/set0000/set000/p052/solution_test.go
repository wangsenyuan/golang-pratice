package p052

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := totalNQueens(n)

	if res != expect {
		t.Errorf("Sample %d expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 92)
}
