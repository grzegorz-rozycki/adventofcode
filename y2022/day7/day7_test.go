package day7

import (
	"adventofcode"
	"sort"
	"strconv"
	"strings"
	"testing"
)

const (
	fileSystemTotal = 70_000_000
	requiredSize    = 30_000_000
	sizeThreshold   = 100_000
)

type node interface {
	GetSize() int
}

type file struct {
	size int
}

func (f *file) GetSize() int {
	return f.size
}

type directory struct {
	parent   *directory
	children map[string]node
}

func newDirectory(parent *directory) *directory {
	return &directory{
		parent:   parent,
		children: make(map[string]node),
	}
}

func (d *directory) GetSize() int {
	size := 0
	for _, child := range d.children {
		size += child.GetSize()
	}
	return size
}

func buildDirectoryTree(commands []string) *directory {
	rootDir := newDirectory(nil)
	currentDir := rootDir

	for _, command := range commands {
		if command == "$ cd /" {
			currentDir = rootDir
		} else if command == "$ cd .." {
			currentDir = currentDir.parent
		} else if strings.HasPrefix(command, "$ cd ") {
			dirName := command[5:]
			dir, ok := currentDir.children[dirName]

			if !ok {
				dir = newDirectory(currentDir)
				currentDir.children[dirName] = dir
			}

			currentDir, ok = dir.(*directory)
		} else if command == "$ ls" {
			// do nothing; if we do a ls for a given directory only once this will be ok,
			// else we'll need to reset the size here
		} else {
			// output of ls
			if strings.HasPrefix(command, "dir") {
				continue
			}
			parts := strings.Split(command, " ")

			if len(parts) != 2 {
				continue
			}

			fileSize, _ := strconv.Atoi(parts[0])
			fileName := parts[1]
			currentDir.children[fileName] = &file{fileSize}
		}
	}

	return rootDir
}

func getTotalSizes(dir *directory) []int {
	sizes := []int{}

	for _, child := range dir.children {
		childDir, isDir := child.(*directory)

		if isDir {
			sizes = append(sizes, getTotalSizes(childDir)...)
		}
	}

	return append(sizes, dir.GetSize())
}

func getTotalSizeOfDirectoriesAboveThreshold(rootDir *directory) int {
	// traverse directory tree in depth first manner;
	// we could split it up into separate methods for traversing and a node visitor,
	// but that's an overkill IMO...
	totalSize := 0

	for _, size := range getTotalSizes(rootDir) {
		if size <= sizeThreshold {
			totalSize += size
		}
	}

	return totalSize
}

func getMinDirSizeRequiredForUpdate(rootDir *directory) int {
	free := fileSystemTotal - rootDir.GetSize()
	missing := requiredSize - free
	sizes := getTotalSizes(rootDir)
	sort.Ints(sizes)

	for _, size := range sizes {
		if size >= missing {
			return size
		}
	}

	return 0
}

func TestExample(t *testing.T) {
	var commands = []string{
		"$ cd /",
		"$ ls",
		"directory a",
		"14848514 b.txt",
		"8504156 c.dat",
		"directory d",
		"$ cd a",
		"$ ls",
		"directory e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}
	rootDir := buildDirectoryTree(commands)

	if got, expected := getTotalSizeOfDirectoriesAboveThreshold(rootDir), 95_437; got != expected {
		t.Fatalf("expected result %v got %v", expected, got)
	}

	if expected, got := 48_381_165, rootDir.GetSize(); got != expected {
		t.Fatalf("expected size of %v got %v", expected, got)
	}
}

func TestInput(t *testing.T) {
	commands, err := adventofcode.ReadInputFile("./input.txt")

	if err != nil {
		t.Fatal(err)
	}

	rootDir := buildDirectoryTree(commands)

	if size := getTotalSizeOfDirectoriesAboveThreshold(rootDir); size <= 0 {
		t.Fatal("expected start to be greater than 0")
	} else {
		t.Logf("got size of %v", size)
	}

	if got := getMinDirSizeRequiredForUpdate(rootDir); got <= 0 {
		t.Fatalf("expected size to be GTE 0")
	} else {
		t.Logf("got missing size of %v", got)
	}
}
