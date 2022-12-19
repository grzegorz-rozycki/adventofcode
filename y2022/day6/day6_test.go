package day6

import (
	"adventofcode"
	"testing"
)

const (
	packetLength  = 4
	messageLength = 14
)

var exampleInput = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}
var expectedStartOfPacket = []int{7, 5, 6, 10, 11}
var expectedStartOfMessage = []int{19, 23, 23, 29, 26}

func findFirstUniqueSequence(length int, input string) int {
search:
	for i := 0; i <= len(input)-length; i++ {
		occurrences := make(map[uint8]int)
		slice := input[i : i+length]

		for j := 0; j < length; j++ {
			char := slice[j]
			if _, ok := occurrences[char]; ok {
				// we have a hit
				continue search
			}
			occurrences[char] = 1
		}
		// seems no repetitions
		return i + length
	}
	return -1
}

func TestExample(t *testing.T) {
	for i, line := range exampleInput {
		if expected, got := expectedStartOfPacket[i], findFirstUniqueSequence(packetLength, line); got != expected {
			t.Fatalf("start-of-packet#%d: excepted %v got %v", i, expected, got)
		}
		if expected, got := expectedStartOfMessage[i], findFirstUniqueSequence(messageLength, line); got != expected {
			t.Fatalf("start-of-message#%d: excepted %v got %v", i, expected, got)
		}
	}
}

func TestInput(t *testing.T) {
	lines, err := adventofcode.ReadInputFile("./input.txt")

	if err != nil {
		t.Fatal(err)
	}

	if start := findFirstUniqueSequence(packetLength, lines[0]); start < 0 {
		t.Fatal("start-of-packet: Expected start to be equal or greater than 0")
	} else {
		t.Logf("start-of-packet: Found start at %v", start)
	}

	if start := findFirstUniqueSequence(messageLength, lines[0]); start < 0 {
		t.Fatal("start-of-message: Expected start to be equal or greater than 0")
	} else {
		t.Logf("start-of-message: Found start at %v", start)
	}
}
