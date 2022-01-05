package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := getPaths(input)
	exp := 19
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := getPaths(input)
	exp := 900
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
