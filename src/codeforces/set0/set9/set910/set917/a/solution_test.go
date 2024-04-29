package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "((?))"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "??()??"
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "))((()(()((((()))())()())((())())(((()()(())))))((())()()(()()(())()))()()(()()()(((()(()(()(()))))("
	expect := 88
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "????????????????????????????????????????????????????????????????????????????????????????????????????"
	expect := 2500
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "()()((?(()(((()()(())(((()((())))(()))(()(((((())))()))(((()()()))))))(((((()))))))))"
	expect := 62
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "?????)(???"
	expect := 21
	runSample(t, s, expect)
}
