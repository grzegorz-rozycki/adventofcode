package day1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getMaxElfCalories(name string) (int, error) {
	file, err := os.Open(name)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	max, current := 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			max = Max(max, current)
			current = 0
		} else if num, err := strconv.Atoi(line); err == nil {
			current += num
		}
	}

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	return max, nil
}

func getTop3ElfCalories(name string) ([3]int, error) {
	file, err := os.Open(name)
	max, current := [3]int{}, 0

	if err != nil {
		return max, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			for i := 0; i < 3; i++ {
				if max[i] < current {
					max[i] = current
					sort.Ints(max[:])
					break
				}
			}
			current = 0
		} else if num, err := strconv.Atoi(line); err == nil {
			current += num
		}
	}

	if err = scanner.Err(); err != nil {
		return max, err
	}

	return max, nil
}
