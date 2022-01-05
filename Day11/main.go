package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type element struct {
	value      int
	hasFlashed bool
}

const GRID_SIZE = 10

func main() {
	part1 := part1("input.txt")
	//part2 := part2("input.txt")
	fmt.Printf("Flashes  %v \n", part1)
}

func part1(s string) int {
	const STEPS = 100
	flashes := 0
	elements := readInFileToIntArray(s)

	for step := 0; step < STEPS; step++ {
		for i := 0; i < GRID_SIZE; i++ {
			for j := 0; j < GRID_SIZE; j++ {
				el := elements[i][j]
				el.value += 1
				if el.value > 9 && !el.hasFlashed {
					el.hasFlashed = true
					flashes += 1
					flashes += proccessFlash(i, j, elements)
				}
			}
		}

		resetFlashes(elements)
	}

	return flashes
}

func resetFlashes(elements [][]*element) {
	for _, elementList := range elements {
		for _, element := range elementList {
			if element.hasFlashed {
				element.value = 0
			}
			element.hasFlashed = false
		}
	}
}

func proccessFlash(i, j int, elements [][]*element) int {
	topLeft := flashCords(i-1, j-1, elements)
	top := flashCords(i, j-1, elements)
	topRight := flashCords(i+1, j-1, elements)
	right := flashCords(i+1, j, elements)
	bottomRight := flashCords(i+1, j+1, elements)
	bottom := flashCords(i, j+1, elements)
	bottomLeft := flashCords(i-1, j+1, elements)
	left := flashCords(i-1, j, elements)

	totalFlash := topLeft + top + topRight + right + bottomRight + bottom + bottomLeft + left
	return totalFlash

}

func flashCords(i, j int, elements [][]*element) int {
	if i >= 0 && j >= 0 && i < GRID_SIZE && j < GRID_SIZE {
		el := elements[i][j]
		el.value += 1
		if el.value > 9 && !el.hasFlashed {
			el.hasFlashed = true
			return 1 + proccessFlash(i, j, elements)
		}
	}
	return 0
}

func readInFileToIntArray(name string) [][]*element {
	output := [][]*element{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		linevals := []*element{}
		for _, char := range line {
			el := element{value: int(char - '0'), hasFlashed: false}
			linevals = append(linevals, &el)
		}
		output = append(output, linevals)
	}
	return output
}
