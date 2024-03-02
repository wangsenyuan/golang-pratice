package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaabc"
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "acb"
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "bac"
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "zbca"
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "czddeneeeemigec"
	expect := 6
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "cdefmopqsvxzz"
	expect := 0
	runSample(t, s, expect)
}
