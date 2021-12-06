package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := poinstWithTwoOrMoreNoDiag(input)
	exp := 5
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := poinstWithTwoOrMore(input)
	exp := 12
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
