package day8

import (
	"adventofcode"
	"sort"
	"testing"
)

type treeGrid [][]uint8

func (g treeGrid) Cols() int {
	return len(g[0])
}

func (g treeGrid) Rows() int {
	return len(g)
}

func isVisibleFromAbove(grid *treeGrid, col, row int) bool {
	for i := row - 1; i >= 0; i-- {
		if (*grid)[row][col] <= (*grid)[i][col] {
			return false
		}
	}
	return true
}

func isVisibleFromBelow(grid *treeGrid, col, row int) bool {
	for i := row + 1; i < grid.Rows(); i++ {
		if (*grid)[row][col] <= (*grid)[i][col] {
			return false
		}
	}
	return true
}

func isVisibleFromLeft(grid *treeGrid, col, row int) bool {
	for i := col - 1; i >= 0; i-- {
		if (*grid)[row][col] <= (*grid)[row][i] {
			return false
		}
	}
	return true
}

func isVisibleFromRight(grid *treeGrid, col, row int) bool {
	for i := col + 1; i < grid.Cols(); i++ {
		if (*grid)[row][col] <= (*grid)[row][i] {
			return false
		}
	}
	return true
}

func isVisible(grid *treeGrid, col, row int) bool {
	cols, rows := grid.Cols(), grid.Rows()
	// if on edge it's always visible
	if col == 0 || row == 0 || col == cols-1 || row == rows-1 {
		return true
	}

	return isVisibleFromBelow(grid, col, row) ||
		isVisibleFromAbove(grid, col, row) ||
		isVisibleFromLeft(grid, col, row) ||
		isVisibleFromRight(grid, col, row)
}

func buildTreeGrid(input []string) *treeGrid {
	// setup grid
	cols, rows := len(input[0]), len(input)
	grid := make(treeGrid, rows)

	for row := 0; row < rows; row++ {
		grid[row] = make([]uint8, cols)
	}

	// read input into grid
	for row, line := range input {
		runes := []rune(line)

		for col := 0; col < cols; col++ {
			grid[row][col] = uint8(runes[col]) - 48
		}
	}

	return &grid
}

func countVisibleTrees(grid *treeGrid) int {
	// count visible; edges are always visible, so we can skip check
	visible := 0

	for row := 0; row < grid.Rows(); row++ {
		for col := 0; col < grid.Cols(); col++ {
			if isVisible(grid, col, row) {
				visible += 1
			}
		}
	}

	return visible
}

func getTreeScore(grid *treeGrid, row int, col int) int {
	above, below, left, right, thisHeight := 0, 0, 0, 0, (*grid)[row][col]

	for i := row - 1; i >= 0; i-- {
		above++
		if thisHeight <= (*grid)[i][col] {
			break
		}
	}

	for i := row + 1; i < grid.Rows(); i++ {
		below++
		if thisHeight <= (*grid)[i][col] {
			break
		}
	}

	for i := col - 1; i >= 0; i-- {
		left++
		if thisHeight <= (*grid)[row][i] {
			break
		}
	}

	for i := col + 1; i < grid.Cols(); i++ {
		right++
		if thisHeight <= (*grid)[row][i] {
			break
		}
	}

	return above * below * left * right
}

func getMaxScore(grid *treeGrid) int {
	scores := make([]int, grid.Cols()*grid.Rows())

	for row := 0; row < grid.Rows(); row++ {
		for col := 0; col < grid.Cols(); col++ {
			i := row*grid.Cols() + col
			scores[i] = getTreeScore(grid, row, col)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	return scores[0]
}

func TestExample(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	grid := buildTreeGrid(input)

	if expected, got := 21, countVisibleTrees(grid); got != expected {
		t.Fatalf("expected %v got %v", expected, got)
	}

	if expected, got := 4, getTreeScore(grid, 1, 2); got != expected {
		t.Fatalf("expected score %v got %v", expected, got)
	}

	if expected, got := 8, getTreeScore(grid, 3, 2); got != expected {
		t.Fatalf("expected score %v got %v", expected, got)
	}
}

func TestInput(t *testing.T) {
	input, err := adventofcode.ReadInputFile("./input.txt")

	if err != nil {
		t.Fatal(err)
	}

	grid := buildTreeGrid(input)

	if visibleTrees := countVisibleTrees(grid); visibleTrees <= 0 {
		t.Fatal("expected start to be greater than 0")
	} else {
		t.Logf("got visibleTrees of %v", visibleTrees)
	}

	if treeScore := getMaxScore(grid); treeScore <= 0 {
		t.Fatal("expected treeScore to be greater than 0")
	} else {
		t.Logf("got max treeScore of %v", treeScore)
	}
}
