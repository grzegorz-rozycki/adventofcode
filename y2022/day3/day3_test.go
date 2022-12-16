package day3

import (
	"adventofcode"
	"testing"
)

var lines = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestBackPackScore(t *testing.T) {
	score := 0
	for _, line := range lines {
		score += scoreBackPack(line)
	}

	if expectedScore := 157; score != expectedScore {
		t.Fatalf("expected score of %v got %v", expectedScore, score)
	}
}

func TestInput(t *testing.T) {
	if score, err := adventofcode.ScoreInput("./input.txt", scoreBackPack); err != nil {
		t.Fatal(err)
	} else if score == 0 {
		t.Fatal("expected score to not equal 0")
	} else {
		t.Logf("Got score of %v", score)
	}
}

func TestBackPackBadges(t *testing.T) {
	if actual, expected := scoreBadges(lines), 70; actual != expected {
		t.Fatalf("expected score of %v got %v", expected, actual)
	}

	lines, err := adventofcode.ReadInputFile("./input.txt")

	if err != nil {
		t.Fatal(err)
	} else if score := scoreBadges(lines); score == 0 {
		t.Fatal("expected score to not equal 0")
	} else {
		t.Logf("Got score of %v", score)
	}
}
