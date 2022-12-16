package day2

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// scores
const (
	lose = 0
	draw = 3
	win  = 6
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
)

// Column   | 1 | 2
// -------- | - | -
// Rock     | A | X
// Paper    | B | Y
// Scissors | C | Z
// R > S
// P > R
// S > P

func scoreInput(name string, scoringFunc func(string, string) int) (int, error) {
	file, err := os.Open(name)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}
		totalScore += scoringFunc(parts[0], parts[1])
	}

	return totalScore, nil
}

// Part 1
func scoreMatch1(opponent, you string) int {
	var result, opponentShape, yourShape int

	if opponent == "A" {
		opponentShape = rock
	} else if opponent == "B" {
		opponentShape = paper
	} else {
		opponentShape = scissors
	}

	if you == "X" {
		yourShape = rock
	} else if you == "Y" {
		yourShape = paper
	} else {
		yourShape = scissors
	}

	switch {
	case yourShape == opponentShape:
		result = draw
	case yourShape == rock && opponentShape == scissors:
		fallthrough
	case yourShape == paper && opponentShape == rock:
		fallthrough
	case yourShape == scissors && opponentShape == paper:
		result = win
	default:
		result = lose

	}

	return result + yourShape
}

func scoreInput1(name string) (int, error) {
	return scoreInput(name, scoreMatch1)
}

// part 2
func scoreMatch2(opponentInput, resultInput string) int {
	var opponentShape, yourShape, result int

	if opponentInput == "A" {
		opponentShape = rock
	} else if opponentInput == "B" {
		opponentShape = paper
	} else {
		opponentShape = scissors
	}

	if resultInput == "X" {
		result = lose
	} else if resultInput == "Y" {
		result = draw
	} else {
		result = win
	}

	switch {
	case result == draw:
		yourShape = opponentShape
	case opponentShape == scissors:
		if result == win {
			yourShape = rock
		} else {
			yourShape = paper
		}
	case opponentShape == rock:
		if result == win {
			yourShape = paper
		} else {
			yourShape = scissors
		}
	case opponentShape == paper:
		if result == win {
			yourShape = scissors
		} else {
			yourShape = rock
		}
	default:
		panic(errors.New("unhandled switch branch"))
	}

	return result + yourShape
}

func scoreInput2(name string) (int, error) {
	return scoreInput(name, scoreMatch2)
}
