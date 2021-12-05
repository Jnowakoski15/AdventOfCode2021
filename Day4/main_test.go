package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := finalScore(input)
	exp := 4512
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := lastWinningBoardScore(input)
	exp := 1924
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
