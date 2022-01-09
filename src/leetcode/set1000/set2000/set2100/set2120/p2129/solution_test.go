package p2129

import "testing"

func runSample(t *testing.T, grid [][]int, h int, w int, expect bool) {
	res := possibleToStamp(grid, h, w)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0},
	}
	h := 4
	w := 3
	expect := true
	runSample(t, grid, h, w, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1},
	}
	h := 2
	w := 2
	expect := false
	runSample(t, grid, h, w, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0, 0}, 
		{0, 0, 0, 0, 0}, 
		{0, 0, 1, 0, 0}, 
		{0, 0, 0, 0, 1}, 
		{0, 0, 0, 1, 1},
	}
	h := 2
	w := 2
	expect := false
	runSample(t, grid, h, w, expect)
}
