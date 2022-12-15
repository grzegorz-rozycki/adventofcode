package day1

import (
	"testing"
)

func TestMaxFromInput(t *testing.T) {
	max, err := getMaxElfCalories("./input.txt")

	if err != nil {
		t.Fatal(err)
	}

	if max == 0 {
		t.Fatal("Max should not be equal to 0")
	}

	t.Logf("Calculated max is %v", max)
}

func TestTop3FromInput(t *testing.T) {
	top, err := getTop3ElfCalories("./input.txt")

	if err != nil {
		t.Fatal(err)
	}

	sum := 0
	for val := range top {
		sum += val
	}

	if sum == 0 {
		t.Fatal("Sum should not be equal to 0")
	}

	t.Logf("Calculated max is %v", top[0]+top[1]+top[2])
}
