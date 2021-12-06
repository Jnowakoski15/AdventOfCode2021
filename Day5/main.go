package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	part1 := poinstWithTwoOrMoreNoDiag("input.txt")
	part2 := poinstWithTwoOrMore("input.txt")

	fmt.Printf("Points with more than 2: %v \n", part1)
	fmt.Printf("Points with 2 or more : %v \n", part2)
}

func poinstWithTwoOrMore(s string) int {
	lines := readInFileToIntArray(s)
	xMax, yMax := getMaxCords(lines)
	plane := buildPlan(xMax+1, yMax+1)

	//Plot lines
	for _, l := range lines {
		if l.x1 == l.x2 {
			plotHor(l, plane)
			continue
		}
		// plot vertical
		if l.y1 == l.y2 {
			plotVert(l, plane)
			continue
		}

		// to the right
		if l.x1 < l.x2 {
			// down to the right
			if l.y1 < l.y2 {
				plotDiagDownRight(l, plane)
				continue
			} else {
				plotDiagUpRight(l, plane)
				continue
			}
			// to the left
		} else {
			if l.y1 < l.y2 {
				plotDiagDownLeft(l, plane)
				continue
			} else {
				plotDiagUpLeft(l, plane)
				continue
			}
		}

	}

	return cordsWithMoreThanTwo(plane)

}

func plotDiagUpLeft(l line, plane [][]int) {
	for i, j := l.x1, l.y1; i >= l.x2 && j >= l.y2; i, j = i-1, j-1 {
		plane[i][j]++
	}
}

func plotDiagDownLeft(l line, plane [][]int) {
	for i, j := l.x1, l.y1; i >= l.x2 && j <= l.y2; i, j = i-1, j+1 {
		plane[i][j]++
	}
}

func plotDiagUpRight(l line, plane [][]int) {
	for i, j := l.x1, l.y1; i <= l.x2 && j >= l.y2; i, j = i+1, j-1 {
		plane[i][j]++
	}
}

func plotDiagDownRight(l line, plane [][]int) {
	for i, j := l.x1, l.y1; i <= l.x2 && j <= l.y2; i, j = i+1, j+1 {
		plane[i][j]++
	}
}

func poinstWithTwoOrMoreNoDiag(s string) int {
	lines := readInFileToIntArray(s)
	cleanLines := removeDiagLines(lines)
	xMax, yMax := getMaxCords(cleanLines)
	plane := buildPlan(xMax+1, yMax+1)

	//Plot lines
	for _, l := range cleanLines {
		// plot Diag
		if l.x1 == l.x2 {
			plotHor(l, plane)
		}
		// plot vertical
		if l.y1 == l.y2 {
			plotVert(l, plane)
		}
	}

	return cordsWithMoreThanTwo(plane)

}

func plotHor(l line, plane [][]int) {
	yMax := l.y2
	yMin := l.y1
	if l.y1 > l.y2 {
		yMax = l.y1
		yMin = l.y2
	}
	for i := yMin; i <= yMax; i++ {
		plane[l.x1][i]++
	}
}

func plotVert(l line, plane [][]int) {
	xMax := l.x2
	xMin := l.x1
	if l.x1 > l.x2 {
		xMax = l.x1
		xMin = l.x2
	}

	for i := xMin; i <= xMax; i++ {
		plane[i][l.y1]++
	}
}

func cordsWithMoreThanTwo(plane [][]int) int {
	output := 0
	for _, arr := range plane {
		for _, el := range arr {
			if el >= 2 {
				output++
			}
		}
	}
	return output
}

func buildPlan(xMax, yMax int) [][]int {
	plane := make([][]int, xMax)

	for i := 0; i < xMax; i++ {
		newCol := make([]int, yMax)
		plane[i] = newCol
	}

	return plane
}

func getMaxCords(cleanLines []line) (int, int) {
	xMax := 0
	yMax := 0
	for _, l := range cleanLines {
		if l.x1 > xMax {
			xMax = l.x1
		}
		if l.x2 > xMax {
			xMax = l.x2
		}
		if l.y1 > yMax {
			yMax = l.y1
		}
		if l.y2 > yMax {
			yMax = l.y2
		}
	}
	return xMax, yMax
}

func removeDiagLines(lines []line) []line {
	output := []line{}

	for _, l := range lines {
		if (l.x1 == l.x2) || (l.y1 == l.y2) {
			output = append(output, l)
		}
	}

	return output
}

func readInFileToIntArray(name string) []line {
	output := []line{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, lineString := range lines {

		points := strings.Split(lineString, " -> ")

		point1Nums := strings.Split(points[0], ",")
		point2Nums := strings.Split(points[1], ",")

		x1, _ := strconv.Atoi(point1Nums[0])
		y1, _ := strconv.Atoi(point1Nums[1])

		x2, _ := strconv.Atoi(point2Nums[0])
		y2, _ := strconv.Atoi(point2Nums[1])

		newCords := line{x1, y1, x2, y2}
		output = append(output, newCords)
	}
	return output
}
