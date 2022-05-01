package main

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 25, 1)
}
