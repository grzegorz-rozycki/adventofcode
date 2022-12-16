package adventofcode

import (
	"bufio"
	"os"
)

func ReadInputFile(name string) ([]string, error) {
	file, err := os.Open(name)
	var lines []string

	if err != nil {
		return lines, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ScoreInput(name string, scoringFunc func(line string) int) (int, error) {
	lines, err := ReadInputFile(name)

	if err != nil {
		return 0, err
	}

	totalScore := 0

	for _, line := range lines {
		totalScore += scoringFunc(line)
	}

	return totalScore, nil
}
