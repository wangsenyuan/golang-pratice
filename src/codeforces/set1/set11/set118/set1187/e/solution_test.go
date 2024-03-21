package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 5},
		{2, 6},
		{1, 4},
		{4, 9},
		{9, 7},
		{9, 8},
	}
	expect := 36
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 14
	runSample(t, n, edges, expect)
}
