package day4

import (
	"adventofcode"
	"testing"
)

var lines = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestOverlappingSections1(t *testing.T) {
	if overlapping := overlappingSections(lines, 1); overlapping != 2 {
		t.Fatalf("expected overlapping lines to equal 2 got %v", overlapping)
	}

	lines, err := adventofcode.ReadInputFile("./input.txt")

	if err != nil {
		t.Fatal(err)
	} else if score := overlappingSections(lines, 1); score == 0 {
		t.Fatal("expected score to not equal 0")
	} else {
		t.Logf("Got score of %v", score)
	}
}

func TestOverlappingSections2(t *testing.T) {
	if overlapping := overlappingSections(lines, 2); overlapping != 4 {
		t.Fatalf("expected overlapping lines to equal 4 got %v", overlapping)
	}

	lines, err := adventofcode.ReadInputFile("./input.txt")

	if err != nil {
		t.Fatal(err)
	} else if score := overlappingSections(lines, 2); score == 0 {
		t.Fatal("expected score to not equal 0")
	} else {
		t.Logf("Got score of %v", score)
	}
}
