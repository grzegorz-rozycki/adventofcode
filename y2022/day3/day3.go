package day3

import (
	"fmt"
)

// Part 1
func scoreBackPack(line string) int {
	items := []rune(line)
	length := len(items)

	if length%2 != 0 {
		panic(fmt.Errorf("expected items count to be even got %v", length))
	}

	commonItem := 0

loop:
	for i := 0; i < length/2; i++ {
		for j := length / 2; j < length; j++ {
			if items[i] == items[j] {
				commonItem = int(items[i])
				break loop
			}
		}
	}

	return scoreItem(commonItem)
}

func scoreItem(commonItem int) int {
	// ASCII a-z map to a score between 1-26
	if commonItem >= 97 && commonItem <= 122 {
		return commonItem - 96
	}
	// ASCII A-Z map to a score between 27-52
	if commonItem >= 65 && commonItem <= 90 {
		return commonItem - 38
	}

	return 0
}

// Part 2
func scoreBadges(lines []string) int {
	totalScore := 0

	for group := 0; group < len(lines); group += 3 {
		commonItem := 0
		elf1Items, elf2Items, elf3Items := lines[group], lines[group+1], lines[group+2]
	group:
		for i1 := 0; i1 < len(elf1Items); i1++ {
			for i2 := 0; i2 < len(elf2Items); i2++ {
				if elf1Items[i1] == elf2Items[i2] {
					for i3 := 0; i3 < len(elf3Items); i3++ {
						if elf2Items[i2] == elf3Items[i3] {
							commonItem = int(elf3Items[i3])
							break group
						}
					}
				}
			}
		}
		totalScore += scoreItem(commonItem)
	}

	return totalScore
}
