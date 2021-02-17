package main

import "testing"

func runSample(t *testing.T, n int, S []string, expect int) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	S := []string{"suf", "mas"}
	expect := 2
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	S := []string{"good", "game", "guys"}
	expect := 0
	runSample(t, n, S, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	S := []string{"hell", "bell", "best", "test"}
	expect := 2
	runSample(t, n, S, expect)
}
