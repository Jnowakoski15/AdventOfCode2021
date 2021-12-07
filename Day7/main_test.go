package main

import "testing"

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := getLeastFuel(input)
	exp := 37
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := getLeastFuel2(input)
	exp := 168
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func TestFuelCost(t *testing.T) {
	const input = 11
	rez := getNewPositionFuelCost(input)
	exp := 66
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
