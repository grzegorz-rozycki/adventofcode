package day5

import (
	"adventofcode"
	"testing"
)

func TestExampleInput(t *testing.T) {
	//     [D]
	// [N] [C]
	// [Z] [M] [P]
	// 1   2   3
	stacks := [][]string{
		/* 1 */ {"Z", "N"},
		/* 2 */ {"M", "C", "D"},
		/* 3 */ {"P"},
	}
	instructions := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	if expected, got := "CMZ", getTopCrates(9000, stacks, instructions); got != expected {
		t.Fatalf("expected %v for crane 9000 got %v", expected, got)
	}

	if expected, got := "MCD", getTopCrates(9001, stacks, instructions); got != expected {
		t.Fatalf("expected %v for crane 9001 got %v", expected, got)
	}
}

func TestFullInput(t *testing.T) {
	//[C]         [S] [H]					TOP
	//[F] [B]     [C] [S]     [W]
	//[B] [W]     [W] [M] [S] [B]
	//[L] [H] [G] [L] [P] [F] [Q]
	//[D] [P] [J] [F] [T] [G] [M] [T]
	//[P] [G] [B] [N] [L] [W] [P] [W] [R]
	//[Z] [V] [W] [J] [J] [C] [T] [S] [C]
	//[S] [N] [F] [G] [W] [B] [H] [F] [N] 	BOTTOM
	//1   2   3   4   5   6   7   8   9

	stacks := [][]string{
		/* 1 */ {"S", "Z", "P", "D", "L", "B", "F", "C"},
		/* 2 */ {"N", "V", "G", "P", "H", "W", "B"},
		/* 3 */ {"F", "W", "B", "J", "G"},
		/* 4 */ {"G", "J", "N", "F", "L", "W", "C", "S"},
		/* 5 */ {"W", "J", "L", "T", "P", "M", "S", "H"},
		/* 6 */ {"B", "C", "W", "G", "F", "S"},
		/* 7 */ {"H", "T", "P", "M", "Q", "B", "W"},
		/* 8 */ {"F", "S", "W", "T"},
		/* 9 */ {"N", "C", "R"},
	}

	lines, err := adventofcode.ReadInputFile("./input.txt")

	if err != nil {
		t.Fatal(err)
	}

	instructions := lines[10:]
	t.Log(instructions[0])

	if got := getTopCrates(9_000, stacks, instructions); got == "" {
		t.Fatal("expected result not to be blank for crane 9000")
	} else {
		t.Logf("got %v for crane 9000", got)
	}

	if got := getTopCrates(9_001, stacks, instructions); got == "" {
		t.Fatal("expected result not to be blank for crane 9001")
	} else {
		t.Logf("got %v for crane 9001", got)
	}
}
