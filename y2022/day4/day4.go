package day4

import (
	"strconv"
	"strings"
)

const part = 1

type section struct {
	begin, end, part int
}

func (this *section) overlaps(other *section) bool {
	if this.part == 1 {
		return this.begin <= other.begin && this.end >= other.end
	}
	return this.begin <= other.begin && this.end >= other.begin
}

func newSectionFromColumn(colum string, part int) section {
	parts := strings.Split(colum, "-")
	begin, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return section{begin, end, part}
}

func overlappingSections(lines []string, part int) int {
	overlapping := 0
	for _, line := range lines {
		cols := strings.Split(line, ",")
		if len(cols) != 2 {
			continue
		}
		section1, section2 := newSectionFromColumn(cols[0], part), newSectionFromColumn(cols[1], part)

		if section1.overlaps(&section2) || section2.overlaps(&section1) {
			overlapping++
		}
	}
	return overlapping
}
