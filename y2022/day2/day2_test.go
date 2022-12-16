package day2

import "testing"

type match struct {
	opponent, you string
}

var matches = []match{
	{"A", "Y"},
	{"B", "X"},
	{"C", "Z"},
}

// part 1
func TestSimple1(t *testing.T) {
	totalScore := 0

	for _, m := range matches {
		totalScore += scoreMatch1(m.opponent, m.you)
	}

	if expectedScore := 15; totalScore != expectedScore {
		t.Fatalf("expected total score to equal %v, got %v", expectedScore, totalScore)
	}
}

func TestInput1(t *testing.T) {
	if score, err := scoreInput1("./input.txt"); err != nil {
		t.Fatal(err)
	} else if score == 0 {
		t.Fatal("expected score to not equal 0")
	} else {
		t.Logf("Got score of %v", score)
	}
}

// part 2
func TestSimple2(t *testing.T) {
	totalScore := 0

	for _, m := range matches {
		totalScore += scoreMatch2(m.opponent, m.you)
	}

	if expectedScore := 12; totalScore != expectedScore {
		t.Fatalf("expected total score to equal %v, got %v", expectedScore, totalScore)
	}
}

func TestInput2(t *testing.T) {
	if score, err := scoreInput2("./input.txt"); err != nil {
		t.Fatal(err)
	} else if score == 0 {
		t.Fatal("expected score to not equal 0")
	} else {
		t.Logf("Got score of %v", score)
	}
}
