package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := sumOfLowPoints(input)
	exp := 15
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := sumOfLowPoints(input)
	exp := 900
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
