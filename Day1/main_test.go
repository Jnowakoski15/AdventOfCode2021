package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := depthMesaureIncrease(input)
	exp := 7
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := slideSumIncrease(input)
	exp := 5
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
