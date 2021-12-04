package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := powerConsumption(input)
	exp := 198
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := lifeSupportRating(input)
	exp := 230
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
