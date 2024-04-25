package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abbaa"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "baaaa"
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "agaa"
	expect := 3
	runSample(t, s, expect)
}
