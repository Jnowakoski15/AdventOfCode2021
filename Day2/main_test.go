package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := finalCordMult(input)
	exp := 150
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := finalCordMult2(input)
	exp := 900
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
