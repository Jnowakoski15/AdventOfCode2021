package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := lanterFishSimulation2(input, 80)
	exp := 5934
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := lanterFishSimulation2(input, 256)
	exp := 26984457539
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
