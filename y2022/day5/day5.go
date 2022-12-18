package day5

import (
	"regexp"
	"strconv"
)

type Stack struct {
	values []string
	top    int
}

func NewStack() *Stack {
	return &Stack{make([]string, 1024), -1}
}

func NewStackFromValues(values []string) *Stack {
	s := NewStack()
	for _, v := range values {
		s.Push(v)
	}
	return s
}

func (s *Stack) Push(crate string) {
	s.top++
	s.values[s.top] = crate
}

func (s *Stack) Pop() string {
	val := s.values[s.top]
	s.top--
	return val
}

var re, _ = regexp.Compile(`\d+`)

func parseInstruction(line string) (int, int, int) {
	matches := re.FindAllString(line, -1)

	if len(matches) < 3 {
		return 0, 0, 0
	}

	move, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])

	return move, from - 1, to - 1

}

func getTopCrates(crane int, state [][]string, instructions []string) string {
	stacks := make([]*Stack, len(state))
	for i := 0; i < len(stacks); i++ {
		stacks[i] = NewStackFromValues(state[i])
	}

	for _, instruction := range instructions {
		move, from, to := parseInstruction(instruction)

		if crane == 9001 {
			// this model can move multiple crates at once,
			// so the order doesn't change
			tempStack := NewStack()
			for i := 0; i < move; i++ {
				tempStack.Push(stacks[from].Pop())
			}
			for i := 0; i < move; i++ {
				stacks[to].Push(tempStack.Pop())
			}
		} else {
			for i := 0; i < move; i++ {
				stacks[to].Push(stacks[from].Pop())
			}
		}
	}

	topCrates := ""
	for _, s := range stacks {
		topCrates += s.Pop()
	}

	return topCrates
}
