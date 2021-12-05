package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type values struct {
	number int
	isHit  bool
}

type bingoBoard struct {
	// Bingo boards are only 5x5
	numbers [][]values
}

func main() {
	part1 := finalScore("input.txt")
	part2 := lastWinningBoardScore("input.txt")
	fmt.Printf("Final Score: %v \n", part1)
	fmt.Printf("Last Final Support: %v \n", part2)
}

func lastWinningBoardScore(s string) int {
	callNums, boards := readInFileToIntArrays(s)
	numberOfBoards := len(boards)
	winningBoards := make(map[int]struct{})
	for _, calledNum := range callNums {
		for boardNum, board := range boards {
			_, isAlreadyWinner := winningBoards[boardNum]
			if !isAlreadyWinner {
				isWon := addAndCheckBoard(&board, calledNum)
				if isWon {
					winningBoards[boardNum] = struct{}{}
					if len(winningBoards) == numberOfBoards {
						fmt.Printf("Last Board: %v won with number %v \n", boardNum, calledNum)
						return calculateWinningNumbers(boards[boardNum], calledNum)
					}

				}
			}
		}
	}
	return -1
}

func finalScore(s string) int {
	callNums, boards := readInFileToIntArrays(s)
	isWon := false
	winningBoard := -1
	finalNum := -1
	for _, calledNum := range callNums {
		for boardNum, board := range boards {
			isWon = addAndCheckBoard(&board, calledNum)
			if isWon {
				finalNum = calledNum
				winningBoard = boardNum
				fmt.Printf("Board: %v won with number %v \n", winningBoard, calledNum)
				break
			}
		}
		if isWon {
			break
		}
	}

	return calculateWinningNumbers(boards[winningBoard], finalNum)
}

func calculateWinningNumbers(board bingoBoard, finalNum int) int {
	unusedSum := 0

	for _, xVal := range board.numbers {
		for _, val := range xVal {
			if !val.isHit {
				unusedSum += val.number
			}
		}
	}

	return unusedSum * finalNum
}

func addAndCheckBoard(b *bingoBoard, calledNum int) bool {
	isHit := false
	xCord := 0
	yCord := 0
	xWin := false
	yWin := false

	for x, xVal := range b.numbers {
		for y, val := range xVal {
			if val.number == calledNum {
				//val.isHit = true
				b.numbers[x][y] = values{calledNum, true}
				isHit = true
				xCord = x
				yCord = y
				break
			}
		}
		if isHit {
			break
		}
	}

	if isHit {
		xWin = true
		yWin = true
		// scan x for win
		for i := 0; i < 5; i++ {
			if !b.numbers[xCord][i].isHit {
				xWin = false
				break
			}
		}
		// scan y for win
		for i := 0; i < 5; i++ {
			if !b.numbers[i][yCord].isHit {
				yWin = false
				break
			}
		}

	}
	return xWin || yWin
}

func readInFileToIntArrays(name string) ([]int, []bingoBoard) {
	bingoDrawings := []int{}
	board := []bingoBoard{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	newBoard := bingoBoard{numbers: [][]values{}}
	lineCount := 0
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {

		// get number drawing list
		if i == 0 {
			numberArray := strings.Split(line, ",")
			for _, numStr := range numberArray {
				val, _ := strconv.Atoi(numStr)
				bingoDrawings = append(bingoDrawings, val)
			}
			continue
		}
		// second line always is a new line
		if i == 1 {
			continue
		}

		if line == "" {
			board = append(board, newBoard)
			newBoard = bingoBoard{numbers: [][]values{}}
			lineCount = 0
			continue
		}

		bingoLine := []values{}
		parsedLine := strings.Split(line, " ")

		for _, numberStr := range parsedLine {
			if numberStr == " " || numberStr == "" {
				//skip extra spaces
				continue
			}
			bingoNum, _ := strconv.Atoi(numberStr)
			bingoLine = append(bingoLine, values{bingoNum, false})
		}

		newBoard.numbers = append(newBoard.numbers, bingoLine)
		lineCount++

	}
	//add last bingo board
	board = append(board, newBoard)
	return bingoDrawings, board
}
