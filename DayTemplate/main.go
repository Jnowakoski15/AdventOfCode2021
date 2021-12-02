package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	//part1 := part1("input.txt")
	//part2 := part2("input.txt")
}

func readInFileToIntArray(name string) []int {
	output := []int{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		output = append(output, val)
	}
	return output
}
